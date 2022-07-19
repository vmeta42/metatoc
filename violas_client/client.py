from diem.jsonrpc import Client as DiemClient
from diem import diem_types, utils,  stdlib, ROOT_ADDRESS, TREASURY_ADDRESS
from vtypes.account_state_blob import AccountStateBlobView
from vtypes.local_account import LocalAccount
from diem.diem_types import AccountAddress
from diem.diem_types import ChainId
from vtypes.account_state import AccountState
from vtypes.contants import VLS
from vtypes.transaction import TransactionView
from typing import Optional
from diem.bcs import serialize, deserialize

from move_core_types.hash import new_sha3_256
import time
import vstdlib
import typing



class Client():
    def __init__(self, server_url, root_private):
        self._cli = DiemClient(server_url)
        self._root_account = LocalAccount.from_dict(
            {
                "private_key": root_private,
                "address": ROOT_ADDRESS
            }
        )
        self._treasury_account = LocalAccount.from_dict(
            {
                "private_key": root_private,
                "address": TREASURY_ADDRESS
            }
        )

    def get_account_sequence(self, addr):
        return self._cli.get_account_sequence(addr)

    def meta42_initialize(self, sender: LocalAccount, currency=VLS):
        script = vstdlib.encode_meta42_initialize_script()
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def meta42_accept(self, sender: LocalAccount, currency=VLS):
        script = vstdlib.encode_meta42_accept_script()
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def meta42_mint_token(self, sender: LocalAccount, path: str, currency=VLS):
        if isinstance(path, str):
            path=path.encode("utf-8")
        script = vstdlib.encode_meta42_mint_token_script(path)
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def meta42_share_token_by_path(self, sender: LocalAccount, receiver:AccountAddress, path, metadata=b"", currency=VLS):
        if isinstance(path, str):
            path=path.encode("utf-8")
        token_id = self.from_sha3_256(path)
        return self.meta42_share_token_by_id(sender, receiver, token_id, metadata, currency)

    def meta42_share_token_by_id(self, sender: LocalAccount, receiver:AccountAddress, token_id, metadata=b"", currency=VLS):
        if isinstance(token_id, str):
            token_id=token_id.encode("utf-8")
        script = vstdlib.encode_meta42_share_token_by_id_script(receiver, token_id, metadata)
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def meta42_share_token_by_index(self, sender: LocalAccount, receiver:AccountAddress, index, metadata=b"", currency=VLS):
        script = vstdlib.encode_meta42_share_token_by_index_script(receiver, index, metadata)
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def create_parent_vasp(self, parent_address: AccountAddress, auth_key_prefix, currency, human_name):
        script = stdlib.encode_create_parent_vasp_account_script(
            coin_type=utils.currency_code(currency),
            sliding_nonce=1,
            new_account_address=parent_address,
            auth_key_prefix=auth_key_prefix,
            human_name=human_name,
            add_all_currencies=True
        )
        seq = self._cli.get_account_sequence(self._treasury_account.account_address)
        signed_tx = self._treasury_account.sign(self.create_transaction(self._treasury_account, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def create_child_vasp(self, sender:LocalAccount, child_address: AccountAddress, auth_key_prefix, currency=VLS, child_initial_balance=0):
        script = stdlib.encode_create_child_vasp_account_script(
            coin_type=utils.currency_code(currency),
            child_address=child_address,
            auth_key_prefix=auth_key_prefix,
            add_all_currencies=True,
            child_initial_balance=child_initial_balance,
        )
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def transfer(self, sender: LocalAccount, to_addr:diem_types.AccountAddress , amount: int, currency, metadata: bytes=b""):
        script=stdlib.encode_peer_to_peer_with_metadata_script(
            currency=utils.currency_code(currency),
            payee=to_addr,
            amount=amount,
            metadata=metadata,
            metadata_signature=b"",
        )
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def allow_publishing_module(self, open, currency=VLS):
        script = vstdlib.encode_allow_publishing_module_script(open)
        seq = self._cli.get_account_sequence(self._root_account.account_address)
        signed_tx = self._root_account.sign(self.create_transaction(self._root_account, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def allow_custom_script(self, currency=VLS):
        script = vstdlib.encode_allow_custom_script()
        seq = self._cli.get_account_sequence(self._root_account.account_address)
        signed_tx = self._root_account.sign(self.create_transaction(self._root_account, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def publish_meta42(self, currency=VLS):
        module=vstdlib.encode_meta42_module()
        seq = self._cli.get_account_sequence(self._root_account.account_address)
        signed_tx = self._root_account.sign(self.create_publish_transaction(self._root_account, seq, module, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def publish_compare(self, currency=VLS):
        module=vstdlib.encode_compare_module()
        seq = self._cli.get_account_sequence(self._root_account.account_address)
        signed_tx = self._root_account.sign(self.create_publish_transaction(self._root_account, seq, module, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def get_account(self, addr):
        return self._cli.get_account(addr)

    def get_account_transaction(
        self,
        account_address: typing.Union[diem_types.AccountAddress, str],
        sequence: int
    ) -> Optional[TransactionView]:

        tx = self._cli.get_account_transaction(account_address, sequence, include_events=True)
        return TransactionView(tx)

    def get_transactions(self, start_version, limit):
        txs = self._cli.get_transactions(start_version, limit, include_events=True)
        return [TransactionView(tx) for tx in txs]

    def get_account_state(self, addr) -> Optional[AccountState]:
        ac = self._cli.get_account_state_with_proof(addr)
        blob = AccountStateBlobView.deserialize(bytes.fromhex(ac.blob))
        return AccountState.deserialize(blob.blob)

    def get_paths(self, addr):
        state = self.get_account_state(addr)
        resources = state.get_account_info_resource()
        if resources is None:
            return []
        return [ token.path.decode("utf-8") for token in resources.tokens]

    @staticmethod
    def create_transaction(
        sender: LocalAccount, seq, script, currency
    ) -> diem_types.RawTransaction:
        return diem_types.RawTransaction(
            sender=sender.account_address,
            sequence_number=seq,
            payload=diem_types.TransactionPayload__Script(script),
            max_gas_amount=1_000_000,
            gas_unit_price=0,
            gas_currency_code=currency,
            expiration_timestamp_secs=int(time.time()) + 30,
            chain_id=ChainId.from_int(4),
        )

    @staticmethod
    def create_publish_transaction(
        sender: LocalAccount, seq, module, currency
    ) -> diem_types.RawTransaction:
        return diem_types.RawTransaction(
            sender=sender.account_address,
            sequence_number=seq,
            payload=diem_types.TransactionPayload__Module(module),
            max_gas_amount=1_000_000,
            gas_unit_price=0,
            gas_currency_code=currency,
            expiration_timestamp_secs=int(time.time()) + 30,
            chain_id=ChainId.from_int(4),
        )

    @classmethod
    def from_sha3_256(cls, data):
        data = serialize(data, type(data))
        sha3 = new_sha3_256()
        sha3.update(data)
        return sha3.digest()

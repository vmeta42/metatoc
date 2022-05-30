from diem.jsonrpc import Client as DiemClient
from diem import diem_types, utils,  stdlib
from vtypes.account_state_blob import AccountStateBlobView
from diem.testing.local_account import LocalAccount
from diem.diem_types import AccountAddress
from diem.diem_types import ChainId
from vtypes.account_state import AccountState
from vtypes.contants import VLS, server_url
from vtypes.transaction import TransactionView

from typing import Optional

import time
import vstdlib
import typing


class Client():
    def __init__(self, server_url):
        self._cli = DiemClient(server_url)


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

    def meta42_mint_token(self, sender: LocalAccount, hdfs_path: bytes, currency=VLS):
        if isinstance(hdfs_path, str):
            hdfs_path=hdfs_path.encode("utf-8")
        script = vstdlib.encode_meta42_mint_token_script(hdfs_path)
        seq = self._cli.get_account_sequence(sender.account_address)
        signed_tx = sender.sign(self.create_transaction(sender, seq, script, currency))
        self._cli.submit(signed_tx)
        return self._cli.wait_for_transaction(signed_tx)

    def meta42_share_token_by_id(self, sender: LocalAccount, receiver:AccountAddress, token_id, metadata=b"", currency=VLS):
        if isinstance(token_id, str):
            token_id=bytes.fromhex(token_id)
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

    def create_child_vasp(self, sender:LocalAccount, child_address: AccountAddress, auth_key_prefix, currency, child_initial_balance=0):
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

    def transfer(self, sender: LocalAccount, to_addr:diem_types.AccountAddress , amount: int, currency, metadata: bytes):
        script=stdlib.encode_peer_to_peer_with_metadata(
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
        return [ TransactionView(tx) for tx in txs]

    def get_account_state(self, addr) -> Optional[AccountState]:
        ac = self._cli.get_account_state_with_proof(addr)
        blob = AccountStateBlobView.deserialize(bytes.fromhex(ac.blob))
        return AccountState.deserialize(blob.blob)

    def get_hdfs_paths(self, addr):
        state = self.get_account_state(addr)
        resources = state.get_account_info_resource()
        if resources is None:
            return []
        return [ token.hdfs_path for token in resources.tokens]


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





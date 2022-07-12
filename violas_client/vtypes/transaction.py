from canoser import Struct
from vtypes.contants import hash_to_type_map, Meta42CodeType, name_to_type_map
from vtypes.contants import minted_events_key, shared_events_key
from vtypes.events import MintedTokenEvent, ShareTokenEvent

class TransactionView(Struct):
    def __init__(self, tx):
        self._tx = tx

    def is_executed(self):
        return self._tx.vm_status.type == "executed"

    def get_sender(self):
        return self._tx.transaction.sender

    def get_sequence_number(self):
        return self._tx.transaction.sequence_number

    def get_tx_type(self):
        tx_hash = self._tx.transaction.script_hash
        if tx_hash != "0000000000000000000000000000000000000000000000000000000000000000":
            script_hash = self._tx.transaction.script_hash
            return hash_to_type_map.get(script_hash)
        return name_to_type_map.get(self._tx.script.type)

    def get_receiver(self):
        if self.get_tx_type() in (Meta42CodeType.META42_SHARE_TOKEN_BY_ID, Meta42CodeType.META42_SHARE_TOKEN_BY_INDEX):
            shared_event = self.get_shared_token_event()
            if shared_event:
                return shared_event.receiver

    def get_token_id(self):
        if self.get_tx_type() in (Meta42CodeType.META42_SHARE_TOKEN_BY_ID, Meta42CodeType.META42_SHARE_TOKEN_BY_INDEX):
            shared_event = self.get_shared_token_event()
            if shared_event:
                return shared_event.token_id
        if self.get_tx_type() in (Meta42CodeType.META42_MINT_TOKEN,):
            minted_event = self.get_minted_token_event()
            if minted_event:
                return minted_event.token_id

    def get_path(self):
        if self.get_tx_type() in (Meta42CodeType.META42_MINT_TOKEN,):
            minted_event = self.get_minted_token_event()
            if minted_event:
                return minted_event.path

    def get_minted_token_event(self):
        for event in self._tx.events:
            if event.key == minted_events_key:
                return MintedTokenEvent.deserialize(bytes.fromhex(event.data.bytes))

    def get_shared_token_event(self):
        for event in self._tx.events:
            if event.key == shared_events_key:
                return ShareTokenEvent.deserialize(bytes.fromhex(event.data.bytes))

    def get_script_hash(self):
        return self._tx.transaction.script_hash

    def __str__(self):
        return self._tx.__str__()
import json
from .result_code import ResultCode

class Response():
    def __init__(self):
        self.status = ResultCode.SUCCESSFUL;
        self.data = {}

    def __str__(self) -> str:
        resp = {};
        resp["code"] = self.status.value
        resp["message"] = self.status.name

        if bool(self.data):
            resp["data"] = self.data;

        return json.dumps(resp);

    def setStatus(self, status):
        self.status = status

    def setData(self, data):
        self.data = data

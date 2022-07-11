import json
from ResultCode import ResultCode

class Response():
    def __init__(self):
        self.code = ResultCode.SUCCESSFUL;
        self.message = "ok";
        self.data = {}

    def __str__(self) -> str:
        resp = {};
        resp["code"] = self.code;
        resp["message"] = self.message;
        if bool(self.data):
            resp["data"] = self.data;

        return json.dumps(resp);

    def setCode(self, code):
        self.code = code;

    def setMessage(self, message):
        self.message = message;

    def setData(self, data):
        self.data = data

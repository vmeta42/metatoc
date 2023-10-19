# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from enum import Enum

class ResultCode(Enum):
    # Successful
    SUCCESSFUL = 0

    # Parameter Error
    PARAM_IS_INVALID = 10001
    PARAM_IS_BLANK = 10002
    PARAM_TYPE_ERROR = 10003
    PARAM_NOT_COMPLETE = 10004

    # Premission Error
    PERMISSION_NO_ACCESS = 20001

    # Data Error
    RESULT_DATA_NONE = 30001
    DATA_ALERADY_EXISTED = 30002

    # Interface Error
    INTERFACE_INNER_INVOKE_ERROR = 40001

#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .datetime_utils import fromisoformat



DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)

@dataclass_json
@dataclass
class EquipmentPortInput:
    name: str
    id: Optional[str] = None
    index: Optional[int] = None
    visibleLabel: Optional[str] = None
    portTypeID: Optional[str] = None
    bandwidth: Optional[str] = None

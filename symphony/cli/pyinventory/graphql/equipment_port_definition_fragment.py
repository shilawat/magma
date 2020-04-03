#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

QUERY: str = """
fragment EquipmentPortDefinitionFragment on EquipmentPortDefinition {
  id
  name
  index
  visibleLabel
}

"""

@dataclass
class EquipmentPortDefinitionFragment(DataClassJsonMixin):
    id: str
    name: str
    index: Optional[int]
    visibleLabel: Optional[str]
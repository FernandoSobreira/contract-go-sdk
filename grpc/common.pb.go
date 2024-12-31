syntax = "proto3";

option go_package = "/server";

enum Network {
  BTC_MAIN = 0;
  BTC_TEST = 1;
  ETH_MAIN = 2;
  ETH_GOERLI = 3;
  ETH_SEPOLIA = 4;
  TRON_MAIN = 5;
  TRON_NILE = 6;
  TRON_SHASTA = 7;
  OKX_MAIN = 8;
  OKX_TEST = 9;
}



message GenerateAccountRequest {
  Network network = 1;
  optional string address = 2;
  optional string private_key = 3;
}

message GenerateAccountResponse {
  string address = 1;
  string private_key = 2;
}

message QueryBalanceRequest {
  Network network = 1;
  string address = 2;
}

message QueryBalanceResponse {
  uint64 balance = 1;
}

message QueryBalance20Request {
  Network network = 1;
  string address = 2;
  string contract_address = 3;
}

message QueryBalance20Response {
  uint64 balance = 3;
}

message QueryNowBlockTransRequest {
  Network network = 1;
  optional string block_id = 2;
}

message QueryNowBlockTransResponse {
  string txid = 1;
  uint64 number = 2;
  uint64 fee_limit = 3;
  string from_address = 4;
  string to_address = 5;
  string contract_address = 6;
}

message QueryPendingBlockTransRequest {
  Network network = 1;
  optional string block_id = 2;
}

message QueryPendingBlockTransResponse {
  string txid = 1;
  uint64 number = 2;
  uint64 fee_limit = 3;
  string from_address = 4;
  string to_address = 5;
  string contract_address = 6;
}

message GenerateTransRequest {
  Network network = 1;
  string from_address = 2;
  string from_address_private_key = 3;
  string to_address = 4;
  uint64 number = 5;
  uint64 gas_limit = 6;
}

message GenerateTransResponse {
  string txid = 1;
  uint64 fee_limit = 2;
}

message GenerateTrans20Request {
  Network network = 1;
  string from_address = 2;
  string from_address_private_key = 3;
  string to_address = 4;
  string contract_address = 5;
  uint64 number = 6;
  uint64 gas_limit = 7;
}

message GenerateTrans20Response {
  string txid = 1;
  uint64 fee_limit = 2;
}

message GenerateApprovalTrans20Request {
  Network network = 1;
  string from_address = 2;
  string from_address_private_key = 3;
  string to_address = 4;
  string contract_address = 5;
  uint64 number = 6;
  uint64 gas_limit = 7;
  string approval_address = 8;
}

message GenerateApprovalTrans20Response {
  string txid = 1;
  uint64 fee_limit = 2;
}




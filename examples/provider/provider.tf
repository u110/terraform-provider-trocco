terraform {
  required_providers {
    trocco = {
      source = "registry.terraform.io/primenumber-dev/trocco"
    }
  }
}

variable "trocco_api_key" {
  type      = string
  sensitive = true
}

provider "trocco" {
  api_key = var.trocco_api_key
  region  = "japan"
}

resource "trocco_datamart_definition" "minimum" {
  name                     = "example_minimum"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
}

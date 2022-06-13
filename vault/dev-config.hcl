disable_mlock = true
ui = true
listener "tcp" {
    tls_disable = 1
    address = "[::]:8200"
    cluster_address = "[::]:8201"
}
storage "raft" {
    path = "./vault/data"
}
cluster_addr = "http://localhost:8201/"
api_addr = "http://localhost:8200/"
resource "nexus_security_user" "admin" {
  userid    = "testAdmin"
  firstname = "test"
  lastname  = "Admin"
  email     = "nexus@example.com"
  password  = "admin123"
  roles     = ["nx-admin"]
  status    = "active"
  source    = "default"
}
package config

var defaultConfig = `# EQCP Configuration

# Application Programming Interface - the primary backend service
[API]
	# Set the host the API should listen on. 
	# By default binds to all devices, you can set e.g. "127.0.0.1:8081"
	# default: ":8081"
	host = ":8081"

# Javascript Web Tokens - used for authentication and token generation
[JWT]
	# please use:
	# openssl genrsa -out eqcp.rsa 512
	# to generate a new RSA private key custom to your server
	# default: "eqcp.rsa"
	privateKeyPath = "eqcp.rsa"
	# please use:
	# openssl rsa -in eqcp.rsa -pubout > eqcp.rsa.pub
	# to generate a new RSA public key custom to your server
	# defaut: "eqcp.rsa.pub"
	publicKeyPath = "eqcp.rsa.pub"

# LoginServer exposes services we use for login
[LoginServer]
	webApiHost = "127.0.0.1:6000"
	apiToken = "oijdfoisjdfiosjdf"
	enabled = true

# Google Remote Procedure Calls - API uses this internally
# can be safely ignored and left at default settings
[GRPC]
	# Set the host GRCP should listen on.
	# Unless you are familiar with GRPC and want to tap into it,
	# this should be kept to default
	# default: 127.0.0.1:9090"
	host = "127.0.0.1:9090"

# Permissions for endpoints
[Permissions.Staff]
	status = 20
[Permissions.Staff.Endpoints.Account.Read]
	fields = [ "*" ]
[Permissions.Player]
	status = 0
[Permissions.Player.Endpoints.Account.Read]
	loginNotRequired = true
	selfOnly = true
	fields = [ "id", "name", "timecreation", "expansion" ]
[Permissions.Player.Endpoints.Character.Read]
	fields = [ "*" ]
[Permissions.Player.Endpoints.Character.Search]
	fields = [ "*" ]
	
#can be ignored. eqemu_config.json settings are used by default
#[Database]
#	host = ""
#	port = ""
#	username = ""
#	password = ""
#	db = ""

`

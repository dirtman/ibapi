# NAME

ibapi url - make a get, post, put or delete Infoblox API request

# USAGE

- ibapi url &lt;add|get|update|delete> <URL>

# DESCRIPTION

"ibapi url" makes an Infoblox API request with the specified
method and URL.

For more details, invoke the specific method
with the --help|-h option. For example:

- ibapi url get -h

# EXAMPLES

- ibapi url get '/record:host?name\~=cs.rice.edu'

    Retrieve all Host records with a name that matches the pattern "cs.rice.edu".

- ibapi url get '/record:a?ipv4addr\~=128.42.201.'

    Retrieve all A records with an IP address that matches the pattern "128.42.201.".

- ibapi url get '/record:host?\_schema'

    Retrieve the schema for a Host record.

# FILES

- /usr/site/ibapi-0.0/etc/ibapi.conf
- /etc/opt/ibapi/ibapi.conf
- /etc/opt/ibapi-0.0//ibapi.conf
- ~/.ibapi/ibapi.conf
- ~/.ibapi-0.0/ibapi.conf

    The IBAPI configuration files which can be used to
    set defaults for nearly all of the options described above.
    Any combination of these may be used.
    Each file found is read in turn, with settings in later files
    overriding those in previous files.  Note that command line
    options override all config file settings.

# SEE ALSO

ibapi(1),
host(1),
host:add(1),
ptr(1),
cname(1),
alias(1),
host:get(1),
host:delete(1),
host:update(1),
ptr:add(1),
ptr:delete(1),
ptr:get(1),
ptr:update(1),
cname:add(1),
cname:delete(1),
cname:get(1),
cname:update(1),
alias:add(1),
alias:delete(1),
alias:get(1),
alias:update(1),
url:add(1),
url:delete(1),
url:get(1),
url:update(1),
a(1),
a:delete(1),
a:get(1),
a:update(1),
a:add(1),
ibapi.conf(5)

# NAME

ibapi cname - create, read, update and delete Infoblox CNAME records

# USAGE

- ibapi cname &lt;add|get|update|delete> &lt;options/args>

# DESCRIPTION

"ibapi cname" can be used to add, get, update and delete Infoblox CNAME records.
The basic format is

- ibapi cname &lt;operation> &lt;options/args>

For more details, invoke the specific operation
with the --help|-h option. For example:

- ibapi cname add -h

# EXAMPLES

- ibapi cname add -t 600 rb4.rice.edu somewhere.com

    Add a CNAME record.

- ibapi cname delete rb4.rice.edu somewhere.com

    Delete a CNAME record.  

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
url(1),
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

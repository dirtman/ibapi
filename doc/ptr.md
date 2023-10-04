# NAME

ibapi ptr - create, read, update and delete Infoblox PTR records

# USAGE

- ibapi ptr &lt;add|get|update|delete> &lt;options/args>

# DESCRIPTION

"ibapi ptr" can be used to add, get, update and delete Infoblox PTR records.
The basic format is

- ibapi ptr &lt;operation> &lt;options/args>

For more details, invoke the specific operation
with the --help|-h option. For example:

- ibapi ptr add -h

# EXAMPLES

- ibapi ptr add -t 600 rb4.rice.edu 168.7.56.224

    Add a PTR record.

- ibapi ptr delete rb4.rice.edu 168.7.56.224

    Delete a PTR record.  

# FILES

- /usr/site/ibapi-1.0/etc/ibapi.conf
- /etc/opt/ibapi/ibapi.conf
- /etc/opt/ibapi-1.0//ibapi.conf
- ~/.ibapi/ibapi.conf
- ~/.ibapi-1.0/ibapi.conf

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
fixedaddress(1),
fixedaddress:add(1),
fixedaddress:delete(1),
fixedaddress:get(1),
fixedaddress:update(1),
grid(1),
mx:add(1),
mx:delete(1),
mx:get(1),
mx:update(1),
mx(1),
txt:add(1),
txt:delete(1),
txt:get(1),
txt:update(1),
txt(1),
ibapi.conf(5)

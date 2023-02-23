# NAME

ibapi - create, read, update and delete basic Infoblox records

# USAGE

- ibapi <-h|--help>
- ibapi &lt;host|a|ptr|cname|alias|url> <-h|--help>
- ibapi &lt;host|a|ptr|cname|alias|url> &lt;add|get|update|delete> &lt;options/args>

# DESCRIPTION

ibapi can be used to add, get, update and delete a few of the most basic Infoblox
records.  The basic format is

- ibapi &lt;record\_type|url> &lt;operation> &lt;options/args>

For more details, invoke the specific record\_type/operation
with the --help|-h option. For example:

- ibapi a -h
- ibapi a add -h

# Configuration Files

ibapi configuration files can be used to set defaults for most of the available
options. The ibapi command searches for configuration files in several places,
including:

- /usr/site/ibapi-0.0/etc/ibapi.conf
- /etc/opt/ibapi/ibapi.conf
- ~/.ibapi/ibapi.conf

Any combination of these may be used. Each file found is read in turn, with
settings in later files overriding those in previous files. Settings specified
via the command line override all configuration file settings.

Configuration file format is very basic and is best shown with an example file:

    # ibapi.conf - ibapi configuration file

    # Global options - apply to (and must be valid for) all sub-commands:
    APIBaseURL =  https://infoblox.go.com/wapi/v2.11
    APIAuthMethod = Basic
    APIAuthTokenID = ibapi        ## You may want to change this.

    # Options specific to the "ibapi a add" sub-command:
    [a:add]
    comment = this is my new A record

Spaces before and after the equal sign are discarded.
Inline comments are allowed; the "#" must be preceeded with one or
more spaces, and these spaces are discarded (along with the comment).

A setting name is equal to a command's long option name without the "--"
prefix.  Case is not sensitive.  Use the --showconfig option to see the
current options and values:

- ibapi url add --showconfig

# Authentication

ibapi supports only basic authentication.  A username and password
can be specified via command line options or via a "username:password"
string stored in a local file.  Here is an example of configuring a username and
password for WAPI user "sandman":

- mkdir -p $HOME/.ibapi/private
- chmod 700 $HOME/.ibapi/private
- echo "APIAuthTokenID = sandman" > $HOME/.ibapi/ibapi.conf
- echo "sandman:WAPI\_PASSWORD" > $HOME/.ibapi/private/sandman
- chmod 600 $HOME/.ibapi/private/sandman

# EXAMPLES

- ibapi host add rb3.rice.edu 168.7.56.225 -d -m f4:8e:38:84:89:e6 -N 10.128.81.10 -b /grub2/shim.efi

    Create the specified Host record with IPv4 address 168.7.56.225, configure that address
    for DHCP and set DHCP-related options as specified.

- ibapi host update rb3.rice.edu -i +168.7.56.226

    Update the Host record named "rb3.rice.edu", adding the IPv4 address "168.7.56.226".

- ibapi host update rb4.rice.edu 168.7.56.224 -i 168.7.56.225

    Update the "rb4.rice.edu" Host record, changing the IP address "168.7.56.224"
    to "168.7.56.225".

- ibapi host get rb4.rice.edu

    Fetch and print the Host record named "rb4.rice.edu".

- ibapi host get 168.7.56.224

    Fetch and print all Host records with IPv4 address "168.7.56.224".

- ibapi host get rb4.rice.edu 168.7.56.224

    Fetch and print the Host record with name "rb4.rice.edu" and IPv4 address "168.7.56.224".

- ibapi host get rb3.rice.edu -v -I bootfile,nextserver,mac

    Fetch and print the Host record named "rb3.rice.edu", including in the output
    the specified IPv4 fields.

- ibapi a get -Fipv4addr\~=10.10.10.20,name\~=a -V external

    Fetch all A records in the external view whose IPs match the pattern "10.10.10.20"
    and names matches the pattern "a".

- ibapi a get -Fname\~=seci.rice.edu

    Fetch all A records whose names match the pattern "seci.rice.edu".

- ibapi a get rb4.rice.edu 168.7.56.224

    Fetch the A record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

- ibapi a get 168.7.56.224

    Get all A records that contain the IP address "168.7.56.224".

- ibapi a add rb4.rice.edu 168.7.56.224

    Create a new A record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

- ibapi a delete rb4.rice.edu 168.7.56.224

    Delete the A record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

- ibapi a update rb4.rice.edu 168.7.56.224 -i 168.7.56.225

    Update the "rb4.rice.edu/168.7.56.224" A record, changing the IP address
    to "168.7.56.225".

- ibapi alias add rb4.rice.edu somewhere.com

    Create a new Alias record with hostname "rb4.rice.edu" and A record target "somewhere.com".

- ibapi cname get rb4.rice.edu

    Fetch the CNAME record with hostname "rb4.rice.edu".

- ibapi cname get -Fcanonical=somewhere.com

    Get each CNAME record whose target (canonical) is "somewhere.com".

- ibapi cname get "" somewhere.com

    Same as above.

- ibapi cname get rb4.rice.edu somewhere.com

    Fetch the CNAME record with hostname "rb4.rice.edu" and target "somewhere.com".

- ibapi cname update rb4.rice.edu -t nowhere.com

    Update the "rb4.rice.edu" CNAME record, changing the target to "nowhere.com".

- ibapi ptr add rb4.rice.edu 168.7.56.224

    Create a new PTR record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

- ibapi ptr delete rb4.rice.edu 168.7.56.224

    Delete the PTR record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

- ibapi ptr get rb4.rice.edu 168.7.56.224

    Fetch the PTR record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

- ibapi ptr get 168.7.56.224

    Get all PTR records that contain the IP address "168.7.56.224".

- ibapi ptr update rb4.rice.edu 168.7.56.224 -i 168.7.56.225

    Update the rb4.rice.edu/168.7.56.224 PTR record, changing the IP address
    to "168.7.56.225".

- ibapi url add 'record:a?name=dbx.seci.rice.edu&ipv4addr=10.10.10.201'

    Create a new A record with hostname "dbx.seci.rice.edu" and IP address "10.10.10.201".
    Same as

        ibapi a add dbx.seci.rice.edu 10.10.10.201

- ibapi url delete "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMTAuMTAuMjAx:dbx.seci.rice.edu/default"

    Delete the A record with the specified object reference.

- ibapi cname get dbalias.seci.rice.edu
- ref=$(ibapi cname get -r dbalias.seci.rice.edu | awk '{print $2}')
- ibapi url delete "$ref"

    Get the reference for the dbalias.seci.rice.edu CNAME record, and then
    use the reference to delete the record.

- ibapi url get '/record:host?name\~=cs.rice.edu'

    Retrieve all Host records with a name that matches the pattern "cs.rice.edu".

- ibapi url get '/record:a?ipv4addr\~=128.42.201.'

    Retrieve all A records with an IP address that matches the pattern "128.42.201.".

- ibapi url get '/record:host?\_schema'

    Retrieve the schema for a Host record.

- ibapi url update "record:cname/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMTAuMTAuMjAx:dbx.seci.rice.edu/default?canonical=somewhere.edu"

    Change the target, or canonical name, of the referenced CNAME record.

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
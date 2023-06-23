# NAME

ibapi - Infoblox WAPI command line tool

# USAGE

- ibapi <OBJECT> <OPERATION> \[&lt;args>\]

where OBJECT is one of

> a alias cname fixedaddress grid host url

and OPERATION, for all but the "grid" object, is one of

> add get update delete

# DESCRIPTION

ibapi is a command for adding, reading, updating and deleting basic DNS records
as well as for managing other Infoblox-specific objects via the Infoblox WAPI.
Currently supported object types are the DNS records
A, Alias, CNAME, and PTR;
the Infoblox-specific object types
fixedaddress, grid and host;
and the special type "url", which allows you to manipulate any type of Infoblox object.

Use the ibapi -h/--help option for more details.  For instance:

- ibapi -h
- ibapi host -h
- ibapi host add -h

# Configuration Files

ibapi configuration files can be used to set defaults for most of the available
options. The ibapi command searches for configuration files in several places,
including:

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

ibapi currently supports only basic authentication.  A username and password
can be specified via command line options or via a "username:password"
string stored in a local file.  Here is an example of configuring a username and
password for WAPI user "sandman":

- mkdir -p $HOME/.ibapi/private
- chmod 700 $HOME/.ibapi/private
- echo "APIAuthTokenID = sandman" > $HOME/.ibapi/ibapi.conf
- echo "sandman:WAPI\_PASSWORD" > $HOME/.ibapi/private/sandman
- chmod 600 $HOME/.ibapi/private/sandman

# Building/Installing

ibapi is written in Go and is available at github.com.
Below is an installation example for Fedora/RHEL:

    sudo dnf install golang
    export GOPATH=$HOME/go
    mkdir -p $GOPATH/src
    cd $GOPATH/src
    git clone https://github.com/dirtman/ibapi
    cd ibapi
    go mod tidy
    make
    ./bin/ibapi -h
    make install
    /usr/bin/ibapi -h

# EXAMPLES

- ibapi host add zabbix.rice.edu 168.7.56.225 -d -R -m f4:8e:38:84:89:e6 -N 10.128.95.14 -b "/grub2/grubx64.efi"

    Create the "zabbix.rice.edu" Host record with the specified IPv4 address enabled for DHCP, and when done, issue the "restart\_if\_needed" command to restart Grid services if needed.

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

- ibapi grid restart

    Instruct Infoblox to restart any grid services that need to be restarted,
    generally due to pending updates that require a particular service, such as
    DHCP, is be restarted.

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
fixedaddress(1),
fixedaddress:add(1),
fixedaddress:delete(1),
fixedaddress:get(1),
fixedaddress:update(1),
grid(1),
ibapi.conf(5)

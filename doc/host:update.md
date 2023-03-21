# NAME

ibapi host update - update Infoblox Host records

# USAGE

- ibapi host update &lt;options/args>

# DESCRIPTION

The update command is used to update Infoblox Host records.

To update a specific IP address of a Host record (such as with the --IP or --MAC
option), both the hostname and the specfic IP address to be updated must be
specified, either via the command line or the --filename option.  Otherwise, only
the hostname is required.

# OPTIONS

Some options can be specified with either a short (i.e., -h) or long
(i.e., --help) form.  In the latter case, case is non-sensitive.

Most options have a corresponding configuration file setting
that is equal to the long option name without the "--" prefix.
Command line options always override configuration file settings.
See ibapi.conf(5) for more information.

Note: while the "default" value for an option is sometimes specified in the option's
description, do not rely on this value.  It is always best to use the --ShowConfig
option to verify the real value.

## OPTIONS - General

- -V &lt;view>, --View=&lt;view>:

    Specify the view of the record to update.  Default: "default".

- -D &lt;true|false>, --Disable=&lt;true|false>:

    Update the record's "disable" setting to the specified value.

- -n &lt;new\_hostname>, --Name=&lt;new\_hostname>:

    Update the name of the specified record to "new\_hostname".

- -c &lt;comment>, --Comment=&lt;comment>:

    Update the record's comment.

- --TTL=&lt;ttl>:

    Update the record's TTL.

- -e &lt;true|false>, --enableDNS=&lt;true|false>:

    Update the record's "configure\_for\_dns" setting to the specified value.

- -F &lt;fields>, --Fields=&lt;fields>:

    Specify fields and corresponding values to be updated.  For intance:
    "comment=RT100931",view=default,ttl=900".

- -i \[+|-\]&lt;new\_IP>, --IP=\[+|-\]&lt;new\_IP>:

    If new\_IP is preceeded with a '+', add new\_IP to the Host record's list of IP addresses.
    If new\_IP is preceeded with a '-', remove new\_IP from the Host record's list of IP addresses.
    If neither the above, change the IP address of the specified hostname/IP to "new\_IP".
    Note that in the latter case, both a hostname and an IP address must be specified as
    arguments.

- -C, --Check:

    Before updating a record's hostname or IP address, check if any "related" records
    already exist, and if so do not update the new record.
    Related records are those that share the same hostname and/or IP address, 
    such as a Host record and A record with the same hostname and/or IP address.
    Default: false.

- -f &lt;filename>, --filename=&lt;filename>:

    Specify a filename containing a list of Host records to update.
    Each line must contain a hostname and, depending on the specified options, an IP address.
    Blank lines and lines beginning with "#" are ignored, as is anything on a line
    following a "#".

## OPTIONS - IPv4 related Options

- -d &lt;true|false>, --enableDHCP=&lt;true|false>:

    Update the specified IP address's configure\_for\_dhcp" setting to the specified value.

- -m &lt;newMAC>, --MAC=&lt;newMAC>:

    Update the MAC address of the specified IP address to "newMAC"  Note that with this
    option, both a hostname and an IP address must be specified.

- -b &lt;bootfile>, --bootfile=&lt;bootfile>:

    Update the bootfile of the specified IP address.

- -N &lt;nextserver>, --nextserver=&lt;nextserver>:

    Update the nextserver of the specified IP address.

- -B &lt;bootserver>, --bootserver=&lt;bootserver>:

    Update the bootserver of the specified IP address.

- -I &lt;ipFields>, --ipFields=&lt;ipFields>:

    Specify a comma separated list of field name/value pairs to updated for the
    specified IP address.  For instance,
    bootfile=/grub2/grubx64.efi,nextserver=10.128.81.10.

## OPTIONS - API Options

- --APIBaseURL=&lt;url>:

    API base URL.
    Default: "https://infoblox.rice.edu/wapi/v2.11"

- --HTTPTimeout=&lt;seconds>:

    Timeout in seconds for the HTTP connection.
    Default: 10.

- --APIAuthMethod=&lt;method>:

    WAPI authentication method for accessing the Infoblox API.  Currently,
    only "Basic" authentication (username and password) is supported.

- -u &lt;username>, --username=&lt;username>:

    Specify the username used for basic auth.

- -p &lt;password>, --password=&lt;password>:

    Specify the password used for basic auth.  If this option is specified
    and is non-empty, either the --username option can be used to specify
    the corresponding username, or the current user will be assumed for username.

- -P, --PromptForPassword:

    Prompt the user for the password used for basic authentication.  This is done
    automatically unless one of these options is specified: --password,
    \--APIAuthToken, --APIAuthTokenID.

- --APIAuthToken=&lt;token>:

    As an alternative to specifing a username and password separately, an "authentication token" 
    containing both the username and password, separated by a ':', can be specified.

- --APIAuthTokenID=&lt;tokenID>:

    As an alternative to specifying a username/password or an APIAuthToken
    directly, the name of an authentication token ID which maps to an
    authentication token can be specified.  See also --SecretsDir.
    Default: "owlapi\_basic".

- --SecretsDir=&lt;pathname>:

    Specify a directory where optional "secret files" are kept.  These files can be
    used as a slightly safer alternative to storing authentication credentials in
    the standard configuration files.  The name of a secret file corresponds to a
    "tokenID" configured via --APIAuthTokenID, and the file contains an
    authentication token.  Obviously, these files should be safely guarded or
    avoided altogether.
    Default: "/etc/opt/ibapi/private".

## OPTIONS - Common To All IBAPI Commands

- -h, --help:

    Help; show usage information and exit.

- --showConfig:

    Read in and show all configuration settings and exit.

- -q, --Quiet (--noQuiet):

    Be quieter than normal.

- --Quieter (--noQuieter):

    Quieter mode.  Suppress all messages except warning and error messages.

- -v, --Verbose (--noVerbose):

    Be louder than normal. Over-rides the "--Quiet"  and "-Quieter" options.

- --nopage:

    By default, usage information (see --help) is piped to the pager specified
    by the environment variable "PAGER", if this environment variable is set,
    or by the pager specified by --Pager.
    The --nopage option disables this paging.

- --Pager=&lt;pager>:

    Specify a pager command for paging the usage information (with --help).  By default,
    the environment variable PAGER is used.  If a full path is not specified, the command
    is searched for using the PATH environment variable.

# EXAMPLES

- ibapi host update rb4.rice.edu 168.7.56.224 -i 168.7.56.225

    Update the "rb4.rice.edu" Host record, changing the IP address "168.7.56.224"
    to "168.7.56.225".

- ibapi host update rb4.rice.edu -i +168.7.56.225

    Update the "rb4.rice.edu" Host record, adding the IP address "168.7.56.225"

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

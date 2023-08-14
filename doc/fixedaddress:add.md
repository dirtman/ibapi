# NAME

ibapi fixedaddress add - create Infoblox fixedaddress records

# USAGE

- ibapi fixedaddress add &lt;options/args>

# DESCRIPTION

The add command is used to create Infoblox fixedaddress records.  To create a
single fixedaddress record, a single IPv4 address and a MAC address can be
provided as command line arguments, in either order.  Alternatively, a list of
records to add can be specified in a file (see --filename).

# OPTIONS

Some options can be specified with either a short (i.e., -h) or long
(i.e., --help) form.  In the latter case, case is non-sensitive.

Boolean options (flags) do not require a value.  "-v" is equivalent to "-v=true".
To disable, set to "false" ("-v=false" or "--verbose=false").

Most options have a corresponding configuration file setting
that is equal to the long option name without the "--" prefix.
Command line options always override configuration file settings.
Use the --ShowConfig to view each option and its value.

## OPTIONS - General

- -n &lt;name>, --Name=&lt;name>

    Specify the name of the new record.

- -V &lt;network\_view>, --View=&lt;network\_view>:

    Specify the network\_view for the new record.  Default: "default".

- -D, --Disable:

    Disable the new record.  Default: false.

- -c &lt;comment>, --Comment=&lt;comment>:

    Specify the comment for the new record.
    Alternatively, you can specify this via the --fields option.
    Default: "ibapi:fixedaddress:add".

- -F &lt;fields>, --Fields=&lt;fields>:

    Specify fields and corresponding values for the new record.  For instance:
    "comment=RT100931",network\_view=external".

- -f &lt;filename>, --Filename=&lt;filename>:

    Specify a filename containing a list of fixedaddress records to create.
    Each line should contain an IPv4 address and a MAC address, in either order, separated
    by one or more spaces.
    Blank lines and lines beginning with "#" are ignored, as is anything on a line
    following a "#".

- -R, --restartServices:

    If all record requests are successfully processed, instruct Infoblox to restart
    any grid services that need to be restarted, generally due to pending updates 
    that require a particular service, such as DHCP, is be restarted.

- -b &lt;bootfile>, --bootfile=&lt;bootfile>:

    Specify the bootfile for the specified IP address.

- -N &lt;nextserver>, --nextserver=&lt;nextserver>:

    Specify the nextserver for the specified IP address.

- -B &lt;bootserver>, --bootserver=&lt;bootserver>:

    Specify the bootserver for the specified IP address.

- -I &lt;ipFields>, --ipFields=&lt;ipFields>:

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

- --username=&lt;username>:

    Specify the username used for basic auth.

- --password=&lt;password>:

    Specify the password used for basic auth.  If this option is specified
    and is non-empty, either the --username option can be used to specify
    the corresponding username, or the current user will be assumed for username.

- --PromptForPassword:

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

- --GridReference=&lt;grid\_reference>:

    Specify the Infoblox grid reference ID. This can be used to save a fetch when
    the --restartServices option is specified.  While this option is only relevant
    to a few commands, it is allowed (ignored) by the other commands.

## OPTIONS - Common To All IBAPI Commands

- -h, --help:

    Help; show usage information and exit.

- --showConfig:

    Read in and show all configuration settings and exit.

- -q, --Quiet:

    Be quieter than normal.

- --Quieter:

    Quieter mode.  Suppress all messages except warning and error messages.

- -v, --Verbose:

    Be louder than normal. Over-rides the "--Quiet"  and "-Quieter" options.
    Note such extra details are printed to Stderr so that the normal output
    remains the same regardless of verbosity.

- --page:

    Page help/usage information via the command specified by the --Pager option or
    the environment variable "PAGER".  If neither of these is set, this option 
    is ignored.  Default: true.

- --Pager=&lt;pager>:

    Specify a pager command for paging the usage information (shown with --help).  By default,
    the environment variable PAGER is used.  If a full path is not specified, the command
    is searched for using the PATH environment variable.

# EXAMPLES

- ibapi fixedaddress add 10.143.195.121 c8:1f:66:c1:79:a1 -n zabbix-n1.rice.edu

    Create a new fixedaddress record with IPv4 address "10.143.195.121", MAC address "c8:1f:66:c1:79:a1", and name "zabbix-n1.rice.edu".

- ibapi fixedaddress add 10.10.10.214 -d 64:00:6a:8f:cc:4d -N10.128.81.10 -b/grub2/grubx64.efi -R

    Create a fixedaddress record with the specified IPv4 address, MAC address, nextserver and bootfile.  When done, issue the "restart\_if\_needed" command to restart Grid services if needed.

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
fixedaddress:delete(1),
fixedaddress:get(1),
fixedaddress:update(1),
grid(1),
ibapi.conf(5)

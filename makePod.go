package main

import(
	. "github.com/dirtman/sitepkg"
)

func makePodMap() error {

	PodMap["a"] = `
NAME
    ibapi a - create, read, update and delete Infoblox A records

USAGE
    ibapi a <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi a" can be used to add, get, update and delete Infoblox A records.
    The basic format is

    * ibapi a <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi a add -h

EXAMPLES
    ibapi a add -t 600 rb4.rice.edu 168.7.56.224
        Add an A record.

    ibapi a delete rb4.rice.edu 168.7.56.224
        Delete an A record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["a:add"] = `
NAME
    ibapi a add - create Infoblox A records

USAGE
    ibapi a add <options/args>

DESCRIPTION
    The add command is used to create Infoblox A records. To create a single
    A record, a single hostname and IP address can be provided as command
    line arguments, in either order. Alternatively, a list of records to add
    can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:address:add".

    --TTL=<ttl>:
        Specify the ttl for the new record. Alternatively, you can specify
        this via the --fields option.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        intance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of A records to create. Each
        line should contain a hostname and an IP address, in either order,
        separated by one or more spaces. Blank lines and lines beginning
        with "#" are ignored, as is anything on a line following a "#".

    -C, --Check:
        Before creating a new record, check if any "related" records already
        exist, and if so do not create the new record. Related records are
        those that share the same name and/or content, such as an Host
        record and A record with the same name and/or IP address. Default:
        false.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi a add rb4.rice.edu 168.7.56.224
        Create a new A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), ibapi.conf(5)


`
	PodMap["a:delete"] = `
NAME
    ibapi a delete - delete Infoblox A records

USAGE
    ibapi a delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox A records. To delete a
    single A record, a single hostname and optionally an IP address may be
    provided as command line arguments. Alternatively, a list of records to
    delete can be specified in a file (see --filename).

    If an IP address is specified, the A record to delete must contain that
    IP address, else no address will be deleted. If no IP address is
    specified, the A record is deleted regardless of its IP address(es).
    =head1 OPTIONS

    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of A records to delete. Each
        line should contain a hostname and optionally an IP address, in
        either order, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi a delete rb4.rice.edu 168.7.56.224
        Delete the A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["a:get"] = `
NAME
    ibapi a get - get Infoblox A records

USAGE
    ibapi a get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox A records.

    By default, the hostname and IP address of each fetched record is shown.
    The --verbose option can be specified to print out the raw response from
    the API.

    To fetch a single A record, a single hostname and/or IP address may be
    provided as command line arguments. Alternatively, a list of records to
    fetch can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to fetch. Specify "any" to search for
        records in all views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of records to get. Each line
        should contain a hostname and/or an IP address, in either order,
        separated by one or more spaces. Blank lines and lines beginning
        with "#" are ignored, as is anything on a line following a "#".

    -r <obj_ref>, --Ref=<obj_ref>:
        Instead of showing the name and content of the fetched record(s),
        show the object reference of each record.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi a get rb4.rice.edu 168.7.56.224
        Fetch the A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi a get 168.7.56.224
        Get all A records that contain the IP address "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["alias"] = `
NAME
    ibapi alias - create, read, update and delete Infoblox Alias records

USAGE
    ibapi alias <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi alias" can be used to add, get, update and delete Infoblox Alias
    records. The basic format is

    * ibapi alias <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi alias add -h

EXAMPLES
    ibapi alias add -t 600 rb4.rice.edu somewhere.com
        Add an Alias record.

    ibapi alias delete rb4.rice.edu somewhere.com
        Delete an Alias record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["alias:add"] = `
NAME
    ibapi alias add - create Infoblox Alias records

USAGE
    ibapi alias add <options/args>

DESCRIPTION
    The add command is used to create Infoblox Alias records. To create a
    single Alias record, a single hostname and target can be provided as
    command line arguments. Alternatively, a list of records to add can be
    specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -T <targetType>, --TargetType=<targetType>:
        Specify the target type for the Alias record. Default: A record.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:alias:add".

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        intance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of Alias records to create.
        Each line should contain a hostname and a target, separated by one
        or more spaces. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

    -C, --Check:
        Before creating a new record, check if any "related" records already
        exist, and if so do not create the new record. Related records are
        those that share the same name and/or content. Default: false.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi alias add rb4.rice.edu somewhere.com
        Create a new Alias record with hostname "rb4.rice.edu" and A record
        target "somewhere.com".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["alias:delete"] = `
NAME
    ibapi alias delete - delete Infoblox Alias records

USAGE
    ibapi alias delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox Alias records. To delete a
    single Alias record, a single hostname and optionally a target be
    provided as command line arguments. Alternatively, a list of records to
    delete can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Alias records to delete.
        Each line should contain a hostname and optionally a target,
        separated by one or more spaces. Blank lines and lines beginning
        with "#" are ignored, as is anything on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi alias delete rb4.rice.edu somewhere.com
        Delete the Alias record with hostname "rb4.rice.edu" and target
        "somewhere.com".

    ibapi alias delete rb4.rice.edu
        Delete the "rb4.rice.edu" Alias record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["alias:get"] = `
NAME
    ibapi alias get - get Infoblox Alias records

USAGE
    ibapi alias get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox Alias records.

    By default, the hostname and target of each fetched record is shown. The
    --verbose option can be specified to print out the raw response from the
    API.

    To fetch a single Alias record, a single hostname and/or target may be
    provided as command line arguments. Alternatively, a list of records to
    fetch can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to fetch. Specify "any" to search for
        records in all views. Default: "any".

    -T <targetType>, --TargetType=<targetType>:
        Specify the target type of the Alias record to fetch. Specify "any"
        to search for all target types. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Alias records to get. Each
        line should contain a hostname and/or a target, separated by one or
        more spaces. Blank lines and lines beginning with "#" are ignored,
        as is anything on a line following a "#".

    -r <obj_ref>, --Ref=<obj_ref>:
        Instead of showing the name and content of the fetched record(s),
        show the object reference of each record.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi alias get rb4.rice.edu
        Fetch the Alias record with hostname "rb4.rice.edu".

    ibapi alias get "" somewhere.com
        Get all Alias records that contain the target "somewhere.com".

    ibapi alias get rb4.rice.edu somewhere.com
        Fetch the Alias record with hostname "rb4.rice.edu" and target
        "somewhere.com".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["alias:update"] = `
NAME
    ibapi alias update - update Infoblox Alias records

USAGE
    ibapi alias update <options/args>

DESCRIPTION
    The update command is used to update Infoblox Alias records.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to update. Default: "default".

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -n <new_hostname>, --Name=<new_hostname>:
        Update the name of the specified record to "new_hostname".

    -c <comment>, --Comment=<comment>:
        Update the record's comment.

    --TTL=<ttl>:
        Update the the record's TTL.

    -t <new_target>, --Target=<new_target>:
        Update the target of the specified alias to "new_target".

    -T <new_target_type>, --targetType=<new_target_type>:
        Update the target_type of the specified alias to "new_target_type".

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For intance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Alias records to update.
        Each line must contain a hostname and, depending on the specified
        options, a target. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

    -C, --Check:
        Before updating a record's hostname or target, check if any
        "related" records already exist, and if so do not update the new
        record. Related records are those that share the same hostname
        and/or target, Default: false.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi alias update rb4.rice.edu somewhere.com -t nowhere.com
        Update the "rb4.rice.edu" Alias record, changing the target to
        "nowhere.com".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["a:update"] = `
NAME
    ibapi a update - update Infoblox A records

USAGE
    ibapi a update <options/args>

DESCRIPTION
    The update command is used to update Infoblox A records.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to update. Default: "default".

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -n <new_hostname>, --Name=<new_hostname>:
        Update the name of the specified record to "new_hostname".

    -c <comment>, --Comment=<comment>:
        Update the record's comment.

    --TTL=<ttl>:
        Update the record's TTL.

    -i <new_IP>, --ip=<new_ip>:
        Update the IP of the specified record to "new_IP".

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For intance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of records to update. Each line
        must contain a hostname and, depending on the specified options, an
        IP address. Blank lines and lines beginning with "#" are ignored, as
        is anything on a line following a "#".

    -C, --Check:
        Before updating a record's hostname or IP address, check if any
        "related" records already exist, and if so do not update the new
        record. Related records are those that share the same hostname
        and/or IP address, such as an A record and Host record with the same
        hostname and/or IP address. Default: false.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi a update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the "rb4.rice.edu/168.7.56.224" A record, changing the IP
        address to "168.7.56.225".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:add(1), ibapi.conf(5)


`
	PodMap["cname"] = `
NAME
    ibapi cname - create, read, update and delete Infoblox CNAME records

USAGE
    ibapi cname <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi cname" can be used to add, get, update and delete Infoblox CNAME
    records. The basic format is

    * ibapi cname <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi cname add -h

EXAMPLES
    ibapi cname add -t 600 rb4.rice.edu somewhere.com
        Add a CNAME record.

    ibapi cname delete rb4.rice.edu somewhere.com
        Delete a CNAME record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["cname:add"] = `
NAME
    ibapi cname add - create Infoblox CNAME records

USAGE
    ibapi cname add <options/args>

DESCRIPTION
    The add command is used to create Infoblox CNAME records. To create a
    single CNAME record, a single hostname and target can be provided as
    command line arguments. Alternatively, a list of records to add can be
    specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:cname:add".

    --TTL=<ttl>:
        Specify the ttl for the new record. Alternatively, you can specify
        this via the --fields option.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        intance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of CNAME records to create.
        Each line should contain a hostname and a target, separated by one
        or more spaces. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi cname add rb4.rice.edu somewhere.com
        Create a new CNAME record with hostname "rb4.rice.edu" and target
        "somewhere.com".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:delete(1), cname:get(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["cname:delete"] = `
NAME
    ibapi cname delete - delete Infoblox CNAME records

USAGE
    ibapi cname delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox CNAME records. To delete a
    single CNAME record, a single hostname may be provided as command line
    arguments. Alternatively, a list of records to delete can be specified
    in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of CNAME records to delete.
        Each line should contain a hostname and optionally a target,
        separated by one or more spaces. Blank lines and lines beginning
        with "#" are ignored, as is anything on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi cname delete rb4.rice.edu somewhere.com
        Delete the CNAME record with hostname "rb4.rice.edu" and target
        "somewhere.com".

    ibapi cname delete rb4.rice.edu
        Delete the "rb4.rice.edu" CNAME record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:get(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["cname:get"] = `
NAME
    ibapi cname get - get Infoblox CNAME records

USAGE
    ibapi cname get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox CNAME records.

    By default, the hostname and target (canonical name) of each fetched
    record is shown. The --verbose option can be specified to print out the
    raw response from the API.

    To fetch a single CNAME record, a single hostname and/or target may be
    provided as command line arguments. Alternatively, a list of records to
    fetch can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to fetch. Specify "any" to search for
        records in all views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of CNAME records to get. Each
        line should contain a hostname to be deleted. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

    -r <obj_ref>, --Ref=<obj_ref>:
        Instead of showing the name and content of the fetched record(s),
        show the object reference of each record.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi cname get rb4.rice.edu
        Fetch the CNAME record with hostname "rb4.rice.edu".

    ibapi cname get "" somewhere.com
        Get all CNAME records that contain the target "somewhere.com".

    ibapi cname get rb4.rice.edu somewhere.com
        Fetch the CNAME record with hostname "rb4.rice.edu" and target
        "somewhere.com".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["cname:update"] = `
NAME
    ibapi cname update - update Infoblox CNAME records

USAGE
    ibapi cname update <options/args>

DESCRIPTION
    The update command is used to update Infoblox CNAME records.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to update. Default: "default".

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -n <new_hostname>, --Name=<new_hostname>:
        Update the name of the specified record to "new_hostname".

    -c <comment>, --Comment=<comment>:
        Update the record's comment.

    --TTL=<ttl>:
        Update the the record's TTL.

    -t <new_target>, --Target=<newTarget>:
        Update the target of the specified CNAME record to "newTarget"

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For intance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of CNAME records to update.
        Blank lines and lines beginning with "#" are ignored, as is anything
        on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi cname update rb4.rice.edu -t nowhere.com
        Update the "rb4.rice.edu" CNAME record, changing the target to
        "nowhere.com".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["host"] = `
NAME
    ibapi host - create, read, update and delete Infoblox Host records

USAGE
    ibapi host <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi host" can be used to add, get, update and delete Infoblox Host
    records. The basic format is

    * ibapi host <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi host add -h

EXAMPLES
    ibapi host add -t 600 rb4.rice.edu 168.7.56.224
        Add a Host record.

    ibapi host delete rb4.rice.edu 168.7.56.224
        Delete a Host record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["host:add"] = `
NAME
    ibapi host add - create Infoblox Host records

USAGE
    ibapi host add <options/args>

DESCRIPTION
    The add command is used to create Infoblox Host records. To create a
    single Host record, a single hostname and IP address can be provided as
    command line arguments, in either order. Alternatively, a list of
    records to add can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:host:add".

    --TTL=<ttl>:
        Specify the ttl for the new record. Alternatively, you can specify
        this via the --fields option.

    -e, --EnableDNS:
        Configure the Host record for DNS. Default true.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        intance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of Host records to create. Each
        line should contain a hostname and an IP address, in either order,
        separated by one or more spaces. Blank lines and lines beginning
        with "#" are ignored, as is anything on a line following a "#".

    -C, --Check:
        Before creating a new record, check if any "related" records already
        exist, and if so do not create the new record. Related records are
        those that share the same name and/or content, such as a Host record
        and A record with the same name and/or IP address. Default: false.

  OPTIONS - IPv4 related Options
    -d, --enableDHCP:
        Set the "configure_for_dhcp" flag for the specified IP to true, thus
        enabling the DHCP configuration for the IP address.

    -m <MAC>, --mac=<MAC>:
        Specify the MAC address of the specified IP address.

    -b <bootfile>, --bootfile=<bootfile>:
        Specify the bootfile for the specified IP address.

    -N <nextserver>, --nextserver=<nextserver>:
        Specify the nextserver for the specified IP address.

    -B <bootserver>, --bootserver=<bootserver>:
        Specify the bootserver for the specified IP address.

    -I <ipFields>, --ipFields=<ipFields>:
        Specify a comma separated list of field name/value pairs for the
        specified IP address. For instance,
        bootfile=/grub2/grubx64.efi,nextserver=10.128.81.10.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi host add rb4.rice.edu 168.7.56.224
        Create a new Host record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi host add rb4.rice.edu 10.10.10.214 -d -m 64:00:6a:8f:cc:4d
    -N10.128.81.10 -b/grub2/grubx64.efi
        Create the specified Host with IP address 10.10.10.214, configure
        that IP for DHCP and set DHCP-related options as specified.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["host:delete"] = `
NAME
    ibapi host delete - delete Infoblox Host records

USAGE
    ibapi host delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox Host records. To delete a
    single Host record, a single hostname and optionally an IP address may
    be provided as command line arguments. Alternatively, a list of records
    to delete can be specified in a file (see --filename).

    If an IP address is specified, the Host record to delete must contain
    that IP address, else no host will be deleted. If no IP address is
    specified, the Host record is deleted regardless of its IP address(es).

    Note that deleting a Host record deletes the whole record, along with
    any and all of its IP addresses etc. To delete an IP address from a Host
    record with deleting the Host record itself, use the "host update"
    command.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Host records to delete. Each
        line should contain a hostname and optionally an IP address, in
        either order, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi host delete rb4.rice.edu 168.7.56.224
        Delete the Host record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi host delete rb4.rice.edu
        Delete the "rb4.rice.edu" Host record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["host:get"] = `
NAME
    ibapi host get - get Infoblox Host records

USAGE
    ibapi host get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox Host records.

    By default, the hostname and IP address of each fetched record is shown.
    The --verbose option can be specified to print out the raw response from
    the API.

    To fetch a single Host record, a single hostname and/or IP address may
    be provided as command line arguments. Alternatively, a list of records
    to fetch can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to fetch. Specify "any" to search for
        records in all views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -I <return_ipFields>, --ipFields=<return_ipFields>:
        Specify additional IP fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Host records to get. Each
        line should contain a hostname and/or an IP address, in either
        order, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

    -r <obj_ref>, --Ref=<obj_ref>:
        Instead of showing the name and content of the fetched record(s),
        show the object reference of each record.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi host get rb4.rice.edu 168.7.56.224
        Fetch the Host record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi host get 168.7.56.224
        Get all Host records that contain the IP address "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["host:update"] = `
NAME
    ibapi host update - update Infoblox Host records

USAGE
    ibapi host update <options/args>

DESCRIPTION
    The update command is used to update Infoblox Host records.

    To update a specific IP address of a Host record (such as with the --IP
    or --MAC option), both the hostname and the specfic IP address to be
    updated must be specified, either via the command line or the --filename
    option. Otherwise, only the hostname is required.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to update. Default: "default".

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -n <new_hostname>, --Name=<new_hostname>:
        Update the name of the specified record to "new_hostname".

    -c <comment>, --Comment=<comment>:
        Update the record's comment.

    --TTL=<ttl>:
        Update the record's TTL.

    -e <true|false>, --enableDNS=<true|false>:
        Update the record's "configure_for_dns" setting to the specified
        value.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For intance:
        "comment=RT100931",view=default,ttl=900".

    -i [+|-]<new_IP>, --IP=[+|-]<new_IP>:
        If new_IP is preceeded with a '+', add new_IP to the Host record's
        list of IP addresses. If new_IP is preceeded with a '-', remove
        new_IP from the Host record's list of IP addresses. If neither the
        above, change the IP address of the specified hostname/IP to
        "new_IP". Note that in the latter case, both a hostname and an IP
        address must be specified as arguments.

    -C, --Check:
        Before updating a record's hostname or IP address, check if any
        "related" records already exist, and if so do not update the new
        record. Related records are those that share the same hostname
        and/or IP address, such as a Host record and A record with the same
        hostname and/or IP address. Default: false.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Host records to update. Each
        line must contain a hostname and, depending on the specified
        options, an IP address. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

  OPTIONS - IPv4 related Options
    -d <true|false>, --enableDHCP=<true|false>:
        Update the specified IP address's configure_for_dhcp" setting to the
        specified value. Note this is not a boolean flag - the value "true"
        or "false" must be specified.

    -m <newMAC>, --MAC=<newMAC>:
        Update the MAC address of the specified IP address to "newMAC" Note
        that with this option, both a hostname and an IP address must be
        specified.

    -b <bootfile>, --bootfile=<bootfile>:
        Update the bootfile of the specified IP address.

    -N <nextserver>, --nextserver=<nextserver>:
        Update the nextserver of the specified IP address.

    -B <bootserver>, --bootserver=<bootserver>:
        Update the bootserver of the specified IP address.

    -I <ipFields>, --ipFields=<ipFields>:
        Specify a comma separated list of field name/value pairs to updated
        for the specified IP address. For instance,
        bootfile=/grub2/grubx64.efi,nextserver=10.128.81.10.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi host update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the "rb4.rice.edu" Host record, changing the IP address
        "168.7.56.224" to "168.7.56.225".

    ibapi host update rb4.rice.edu -i +168.7.56.225
        Update the "rb4.rice.edu" Host record, adding the IP address
        "168.7.56.225"

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["ibapi"] = `
NAME
    ibapi - create, read, update and delete basic Infoblox records

USAGE
    ibapi <-h|--help>
    ibapi <host|a|ptr|cname|alias|url> <-h|--help>
    ibapi <host|a|ptr|cname|alias|url> <add|get|update|delete>
    <options/args>

DESCRIPTION
    ibapi can be used to add, get, update and delete a few of the most basic
    Infoblox records. The basic format is

    * ibapi <record_type|url> <operation> <options/args>

    For more details, invoke the specific record_type/operation with the
    --help|-h option. For example:

    * ibapi a -h

    * ibapi a add -h

Configuration Files
    ibapi configuration files can be used to set defaults for most of the
    available options. The ibapi command searches for configuration files in
    several places, including:

    *   /usr/site/ibapi-0.0/etc/ibapi.conf

    *   /etc/opt/ibapi/ibapi.conf

    *   ~/.ibapi/ibapi.conf

    Any combination of these may be used. Each file found is read in turn,
    with settings in later files overriding those in previous files.
    Settings specified via the command line override all configuration file
    settings.

    Configuration file format is very basic and is best shown with an
    example file:

      # ibapi.conf - ibapi configuration file

      # Global options - apply to (and must be valid for) all sub-commands:
      APIBaseURL =  https://infoblox.go.com/wapi/v2.11
      APIAuthMethod = Basic
      APIAuthTokenID = ibapi        ## You may want to change this.

      # Options specific to the "ibapi a add" sub-command:
      [a:add]
      comment = this is my new A record

    Spaces before and after the equal sign are discarded. Inline comments
    are allowed; the "#" must be preceeded with one or more spaces, and
    these spaces are discarded (along with the comment).

    A setting name is equal to a command's long option name without the "--"
    prefix. Case is not sensitive. Use the --showconfig option to see the
    current options and values:

    *   ibapi url add --showconfig

Authentication
    ibapi supports only basic authentication. A username and password can be
    specified via command line options or via a "username:password" string
    stored in a local file. Here is an example of configuring a username and
    password for WAPI user "sandman":

    *   mkdir -p $HOME/.ibapi/private

    *   chmod 700 $HOME/.ibapi/private

    *   echo "APIAuthTokenID = sandman" > $HOME/.ibapi/ibapi.conf

    *   echo "sandman:WAPI_PASSWORD" > $HOME/.ibapi/private/sandman

    *   chmod 600 $HOME/.ibapi/private/sandman

EXAMPLES
    ibapi host add rb3.rice.edu 168.7.56.225 -d -m f4:8e:38:84:89:e6 -N
    10.128.81.10 -b /grub2/shim.efi
        Create the specified Host record with IPv4 address 168.7.56.225,
        configure that address for DHCP and set DHCP-related options as
        specified.

    ibapi host update rb3.rice.edu -i +168.7.56.226
        Update the Host record named "rb3.rice.edu", adding the IPv4 address
        "168.7.56.226".

    ibapi host update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the "rb4.rice.edu" Host record, changing the IP address
        "168.7.56.224" to "168.7.56.225".

    ibapi host get rb4.rice.edu
        Fetch and print the Host record named "rb4.rice.edu".

    ibapi host get 168.7.56.224
        Fetch and print all Host records with IPv4 address "168.7.56.224".

    ibapi host get rb4.rice.edu 168.7.56.224
        Fetch and print the Host record with name "rb4.rice.edu" and IPv4
        address "168.7.56.224".

    ibapi host get rb3.rice.edu -v -I bootfile,nextserver,mac
        Fetch and print the Host record named "rb3.rice.edu", including in
        the output the specified IPv4 fields.

    ibapi a get -Fipv4addr~=10.10.10.20,name~=a -V external
        Fetch all A records in the external view whose IPs match the pattern
        "10.10.10.20" and names matches the pattern "a".

    ibapi a get -Fname~=seci.rice.edu
        Fetch all A records whose names match the pattern "seci.rice.edu".

    ibapi a get rb4.rice.edu 168.7.56.224
        Fetch the A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi a get 168.7.56.224
        Get all A records that contain the IP address "168.7.56.224".

    ibapi a add rb4.rice.edu 168.7.56.224
        Create a new A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi a delete rb4.rice.edu 168.7.56.224
        Delete the A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi a update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the "rb4.rice.edu/168.7.56.224" A record, changing the IP
        address to "168.7.56.225".

    ibapi alias add rb4.rice.edu somewhere.com
        Create a new Alias record with hostname "rb4.rice.edu" and A record
        target "somewhere.com".

    ibapi cname get rb4.rice.edu
        Fetch the CNAME record with hostname "rb4.rice.edu".

    ibapi cname get -Fcanonical=somewhere.com
        Get each CNAME record whose target (canonical) is "somewhere.com".

    ibapi cname get "" somewhere.com
        Same as above.

    ibapi cname get rb4.rice.edu somewhere.com
        Fetch the CNAME record with hostname "rb4.rice.edu" and target
        "somewhere.com".

    ibapi cname update rb4.rice.edu -t nowhere.com
        Update the "rb4.rice.edu" CNAME record, changing the target to
        "nowhere.com".

    ibapi ptr add rb4.rice.edu 168.7.56.224
        Create a new PTR record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi ptr delete rb4.rice.edu 168.7.56.224
        Delete the PTR record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi ptr get rb4.rice.edu 168.7.56.224
        Fetch the PTR record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi ptr get 168.7.56.224
        Get all PTR records that contain the IP address "168.7.56.224".

    ibapi ptr update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the rb4.rice.edu/168.7.56.224 PTR record, changing the IP
        address to "168.7.56.225".

    ibapi url add 'record:a?name=dbx.seci.rice.edu&ipv4addr=10.10.10.201'
        Create a new A record with hostname "dbx.seci.rice.edu" and IP
        address "10.10.10.201". Same as

          ibapi a add dbx.seci.rice.edu 10.10.10.201

    ibapi url delete
    "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMTAuMT
    AuMjAx:dbx.seci.rice.edu/default"
        Delete the A record with the specified object reference.

    ibapi cname get dbalias.seci.rice.edu
    ref=$(ibapi cname get -r dbalias.seci.rice.edu | awk '{print $2}')
    ibapi url delete "$ref"
        Get the reference for the dbalias.seci.rice.edu CNAME record, and
        then use the reference to delete the record.

    ibapi url get '/record:host?name~=cs.rice.edu'
        Retrieve all Host records with a name that matches the pattern
        "cs.rice.edu".

    ibapi url get '/record:a?ipv4addr~=128.42.201.'
        Retrieve all A records with an IP address that matches the pattern
        "128.42.201.".

    ibapi url get '/record:host?_schema'
        Retrieve the schema for a Host record.

    ibapi url update
    "record:cname/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMT
    AuMTAuMjAx:dbx.seci.rice.edu/default?canonical=somewhere.edu"
        Change the target, or canonical name, of the referenced CNAME
        record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["ptr"] = `
NAME
    ibapi ptr - create, read, update and delete Infoblox PTR records

USAGE
    ibapi ptr <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi ptr" can be used to add, get, update and delete Infoblox PTR
    records. The basic format is

    * ibapi ptr <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi ptr add -h

EXAMPLES
    ibapi ptr add -t 600 rb4.rice.edu 168.7.56.224
        Add a PTR record.

    ibapi ptr delete rb4.rice.edu 168.7.56.224
        Delete a PTR record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["ptr:add"] = `
NAME
    ibapi ptr add - create Infoblox PTR records

USAGE
    ibapi ptr add <options/args>

DESCRIPTION
    The add command is used to create Infoblox PTR records. To create a
    single PTR record, a single hostname and IP address can be provided as
    command line arguments, in either order. Alternatively, a list of
    records to add can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:ptr:add".

    --TTL=<ttl>:
        Specify the ttl for the new record. Alternatively, you can specify
        this via the --fields option.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        intance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of PTR records to create. Each
        line should contain a hostname and an IP address, in either order,
        separated by one or more spaces. Blank lines and lines beginning
        with "#" are ignored, as is anything on a line following a "#".

    -C, --Check:
        Before creating a new record, check if any "related" records already
        exist, and if so do not create the new record. Related records are
        those that share the same name and/or content, such as a PTR record
        and A record with the same name and/or IP address. Default: false.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi ptr add rb4.rice.edu 168.7.56.224
        Create a new PTR record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["ptr:delete"] = `
NAME
    ibapi ptr delete - delete Infoblox PTR records

USAGE
    ibapi ptr delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox PTR records. To delete a
    single PTR record, a single hostname and optionally an IP address may be
    provided as command line arguments. Alternatively, a list of records to
    delete can be specified in a file (see --filename).

    If an IP address is specified, the PTR record to delete must contain
    that IP address, else no ptr will be deleted. If no IP address is
    specified, the PTR record is deleted regardless of its IP address(es).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of PTR records to delete. Each
        line should contain a hostname and optionally an IP address, in
        either order, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi ptr delete rb4.rice.edu 168.7.56.224
        Delete the PTR record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:get(1), ptr:update(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["ptr:get"] = `
NAME
    ibapi ptr get - get Infoblox PTR records

USAGE
    ibapi ptr get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox PTR records.

    By default, the name and IP address of each fetched record is shown. The
    --verbose option can be specified to print out the raw response from the
    API.

    To fetch a single PTR record, a single hostname and/or IP address may be
    provided as command line arguments. Alternatively, a list of records to
    fetch can be specified in a file (see --filename).

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view of the record to fetch. Specify "any" to search for
        records in all views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of PTR records to get. Each
        line should contain a hostname and/or an IP address, in either
        order, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

    -r <obj_ref>, --Ref=<obj_ref>:
        Instead of showing the name and content of the fetched record(s),
        show the object reference of each record.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi ptr get rb4.rice.edu 168.7.56.224
        Fetch the PTR record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi ptr get 168.7.56.224
        Get all PTR records that contain the IP address "168.7.56.224".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1),
    url:update(1), a(1), a:delete(1), a:get(1), a:update(1), a:add(1),
    ibapi.conf(5)


`
	PodMap["ptr:update"] = `
NAME
    ibapi ptr update - update Infoblox PTR records

USAGE
    ibapi ptr update <options/args>

DESCRIPTION
    The update command is used to update Infoblox PTR records.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:host:add".

    --TTL=<ttl>:
        Specify the ttl for the new record. Alternatively, you can specify
        this via the --fields option.

    -n <new_hostname>, --Name=<new_hostname>:
        Change the name of the specified ptr to "new_hostname".

    -i <new_IP>, --IP=<new_IP>:
        Change the IP address of the specified ptr to "new_IP".

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For intance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of PTR records to update. Each
        line must contain a hostname and, depending on the specified
        options, an IP address. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

    -C, --Check:
        Before updating a record's hostname or IP address, check if any
        "related" records already exist, and if so do not update the new
        record. Related records are those that share the same hostname
        and/or IP address, such as a PTR record and A record with the same
        hostname and/or IP address. Default: false.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi ptr update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the rb4.rice.edu/168.7.56.224 PTR record, changing the IP
        address to "168.7.56.225".

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), a(1), a:delete(1),
    a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["url"] = `
NAME
    ibapi url - make a get, post, put or delete Infoblox API request

USAGE
    ibapi url <add|get|update|delete> <URL>

DESCRIPTION
    "ibapi url" makes an Infoblox API request with the specified method and
    URL.

    For more details, invoke the specific method with the --help|-h option.
    For example:

    * ibapi url get -h

EXAMPLES
    ibapi url get '/record:host?name~=cs.rice.edu'
        Retrieve all Host records with a name that matches the pattern
        "cs.rice.edu".

    ibapi url get '/record:a?ipv4addr~=128.42.201.'
        Retrieve all A records with an IP address that matches the pattern
        "128.42.201.".

    ibapi url get '/record:host?_schema'
        Retrieve the schema for a Host record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url:add(1), url:delete(1), url:get(1), url:update(1),
    a(1), a:delete(1), a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["url:add"] = `
NAME
    ibapi url add - make an Infoblox POST request

USAGE
    ibapi url add <URL>

DESCRIPTION
    The add command is used to make an Infoblox POST API request.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi url add 'record:a?name=dbx.seci.rice.edu&ipv4addr=10.10.10.201'
        Create a new A record with hostname "dbx.seci.rice.edu" and IP
        address "10.10.10.201". Same as

          ibapi address add dbx.seci.rice.edu 10.10.10.201

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:delete(1), url:get(1), url:update(1), a(1),
    a:delete(1), a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["url:delete"] = `
NAME
NAME
    ibapi url delete - make an Infoblox DELETE request

USAGE
    ibapi url delete <URL>

DESCRIPTION
    The delete command is used to make an Infoblox DELETE API request.

    Note that the Infoblox DELETE method requires an object reference. The
    various ibapi record "get" commands (such as "ibapi address get") all
    have a "--ref" option to get the object reference of an object.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi url delete
    "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMTAuMT
    AuMjAx:dbx.seci.rice.edu/default"
        Delete the A record with the specified object reference.

    ibapi cname get dbalias.seci.rice.edu
    ref=$(ibapi cname get -r dbalias.seci.rice.edu | awk '{print $2}')
    ibapi url delete "$ref"
        Get the reference for the dbalias.seci.rice.edu CNAME record, and
        then use the reference to delete the record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:get(1), url:update(1), a(1),
    a:delete(1), a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["url:get"] = `
NAME
    ibapi url get - make an Infoblox GET request

USAGE
    ibapi url get <URL>

DESCRIPTION
    The get command is used to make an Infoblox GET API request.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - General
    -R <return_fields>, --rFieldsPlus=<return_fields>:
        Specify fields to show in addition to those normally shown.

    -r <return_fields>, --rFields=<return_fields>:
        Specify the fields to show (in addition to the reference (_ref)
        field, which is always shown).

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi url get 'record:host_ipv4addr?ipv4addr=10.143.195.131'
        Retrieve the specified Host ipv4address record.

    ibapi url get
    'record:host_ipv4addr?ipv4addr=10.143.195.131&_return_fields%2b=options,
    bootfile'
        Same as above, but include some additional return fields, if they
        are set.

    ibapi url get '/record:host?name~=cs.rice.edu'
        Retrieve all Host records with a name that matches the pattern
        "cs.rice.edu".

    ibapi url get '/record:a?ipv4addr~=128.42.201.'
        Retrieve all A records with an IP address that matches the pattern
        "128.42.201.".

    url get
    'ipv4address?ip_address=10.143.195.131&_return_fields%2b=extattrs'
        Retrieve the ipv4address record, and add "extattrs" to the list of
        return fields..

    ibapi url get '/record:host?_schema'
        Retrieve the schema for a Host record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:update(1), a(1),
    a:delete(1), a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
	PodMap["url:update"] = `
NAME
    ibapi url update - make an Infoblox PUT request

USAGE
    ibapi url update <URL>

DESCRIPTION
    The update command is used to make an Infoblox PUT API request.

    Note that the Infoblox PUT method requires an object reference. The
    various ibapi record "get" commands (such as "ibapi address get") all
    have a "--ref" option to get the object reference of an object.

OPTIONS
    Some options can be specified with either a short (i.e., -h) or long
    (i.e., --help) form. In the latter case, case is non-sensitive.

    Boolean options (flags) do not require a value. "-v" is equivalent to
    "-v=true". To disable, set to "false" ("-v=false" or "--verbose=false").

    Most options have a corresponding configuration file setting that is
    equal to the long option name without the "--" prefix. Command line
    options always override configuration file settings. Use the
    --ShowConfig to view each option and its value.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    -u <username>, --username=<username>:
        Specify the username used for basic auth.

    -p <password>, --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    -P, --PromptForPassword:
        Prompt the user for the password used for basic authentication. This
        is done automatically unless one of these options is specified:
        --password, --APIAuthToken, --APIAuthTokenID.

    --APIAuthToken=<token>:
        As an alternative to specifing a username and password separately,
        an "authentication token" containing both the username and password,
        separated by a ':', can be specified.

    --APIAuthTokenID=<tokenID>:
        As an alternative to specifying a username/password or an
        APIAuthToken directly, the name of an authentication token ID which
        maps to an authentication token can be specified. See also
        --SecretsDir. Default: "owlapi_basic".

    --SecretsDir=<pathname>:
        Specify a directory where optional "secret files" are kept. These
        files can be used as a slightly safer alternative to storing
        authentication credentials in the standard configuration files. The
        name of a secret file corresponds to a "tokenID" configured via
        --APIAuthTokenID, and the file contains an authentication token.
        Obviously, these files should be safely guarded or avoided
        altogether. Default: "/etc/opt/ibapi/private".

  OPTIONS - Common To All IBAPI Commands
    -h, --help:
        Help; show usage information and exit.

    --showConfig:
        Read in and show all configuration settings and exit.

    -q, --Quiet:
        Be quieter than normal.

    --Quieter:
        Quieter mode. Suppress all messages except warning and error
        messages.

    -v, --Verbose:
        Be louder than normal. Over-rides the "--Quiet" and "-Quieter"
        options. Note such extra details are printed to Stderr so that the
        normal output remains the same regardless of verbosity.

    --page:
        Page help/usage information via the command specified by the --Pager
        option or the environment variable "PAGER". If neither of these is
        set, this option is ignored. Default: true.

    --Pager=<pager>:
        Specify a pager command for paging the usage information (shown with
        --help). By default, the environment variable PAGER is used. If a
        full path is not specified, the command is searched for using the
        PATH environment variable.

EXAMPLES
    ibapi url update
    "record:cname/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMT
    AuMTAuMjAx:dbx.seci.rice.edu/default?canonical=somewhere.edu"
        Change the target, or canonical name, of the referenced CNAME
        record.

    ibapi cname get dbalias.seci.rice.edu
    ref=$(ibapi cname get -r dbalias.seci.rice.edu | awk '{print $2}')
    ibapi url update "${ref}?canonical=somewhere.edu"
        Get the reference for the dbalias.seci.rice.edu CNAME record, and
        then use the reference to update the record.

FILES
    /usr/site/ibapi-0.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-0.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-0.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    ibapi(1), host(1), host:add(1), ptr(1), cname(1), alias(1), host:get(1),
    host:delete(1), host:update(1), ptr:add(1), ptr:delete(1), ptr:get(1),
    ptr:update(1), cname:add(1), cname:delete(1), cname:get(1),
    cname:update(1), alias:add(1), alias:delete(1), alias:get(1),
    alias:update(1), url(1), url:add(1), url:delete(1), url:get(1), a(1),
    a:delete(1), a:get(1), a:update(1), a:add(1), ibapi.conf(5)


`
return nil
}

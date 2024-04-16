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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a:add(1), a:delete(1), a:get(1), a:update(1), alias(1), alias:add(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        instance: "comment=RT100931",view=default,ttl=900".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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

    ibapi a add -V external -f /tmp/hosts
        Create an A record in the "external" view for each hostname/IP pair
        specified in the file "/tmp/hosts".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:delete(1), a:get(1), a:update(1), alias(1), alias:add(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    -V <view>, --View=<view>:
        Specify the view of the record to delete. Default: "default".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:get(1), a:update(1), alias(1), alias:add(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi a get -F name~=mx -V external
        Fetch each A record (in the external DNS view) whose name matches
        "mx".

    ibapi a get -F ipv4addr~=128.42.201.
        Fetch each A record (in any DNS view) whose IP address matches
        "128.42.201.".

    ibapi a get rb4.rice.edu
        Fetch each A record with hostname "rb4.rice.edu".

    ibapi a get 168.7.56.224
        Fetch each A record with IP address "168.7.56.224".

    ibapi a get rb4.rice.edu 168.7.56.224
        Fetch each A record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:update(1), alias(1), alias:add(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        Delete an Alias record with A record target type..

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias:add(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        instance: "comment=RT100931",view=default,ttl=900".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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

    ibapi alias add -T TXT rb4.rice.edu somewhere.com
        Create a new Alias record with hostname "rb4.rice.edu" and TXT
        record target "somewhere.com".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    -V <view>, --View=<view>:
        Specify the view of the record to delete. Default: "default".

    -T <targetType>, --TargetType=<targetType>:
        Specify the target type to delete. Default: A record.

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
        Delete the Alias record with hostname "rb4.rice.edu" and A record
        target "somewhere.com".

    ibapi alias delete rb4.rice.edu -T TXT
        Delete the "rb4.rice.edu" Alias record of target type TXT.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
        Fetch each Alias record with hostname "rb4.rice.edu".

    ibapi alias get "" somewhere.com
        Fetch each Alias record that contains the target "somewhere.com".

    ibapi alias get rb4.rice.edu somewhere.com
        Fetch each Alias record with hostname "rb4.rice.edu" and target
        "somewhere.com".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        Specify fields and corresponding values to be updated. For instance:
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        Specify fields and corresponding values to be updated. For instance:
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), alias(1), alias:add(1),
    alias:delete(1), alias:get(1), alias:update(1), cname(1), cname:add(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["authzone"] = `
NAME
    ibapi authzone - create, read, update and delete Infoblox authoritative
    zones

USAGE
    ibapi authzone <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi authzone" can be used to add, get, update and delete Infoblox
    authoritative zones. The basic format is

    * ibapi authzone <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi authzone add -h

EXAMPLES
    ibapi authzone add t1.zone.rice.edu
        Add a authoritative zone.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["authzone:add"] = `
NAME
    ibapi authzone add - create Infoblox authoritative zones

USAGE
    ibapi authzone add <options/args>

DESCRIPTION
    The add command is used to create Infoblox authoritative zones. To
    create a single authoritative zone, a single zone name (fqdn) can be
    provided as a command line argument. Alternatively, a list of zones to
    add can be specified in a file (see --filename).

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
        Specify the view for the new zone. Default: "default".

    -n <ns_group>, --nsgroup=<ns_group>:
        Specify the name server group for the new zone. By default, if the
        zone is being created in the "external" view, the name server group
        is set to "external_Rice". Otherwise, by default, it is not
        specified (and therefore defaults to the Infoblox default).

    -z <zone_format>, --zoneFormat=<zone_format>:
        Specify the zone format of the zone. Valid values are: FORWARD,
        IPV4, and IPV6. Default: FORWARD.

    -D, --Disable:
        Disable the new zone. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new zone. Alternatively, you can specify
        this via the --fields option. Default: "ibapi:host:add".

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new zone. For
        instance: "comment=RT100931",ns_group="external_Rice".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of authoritative zones to
        create. Each line should contain one zone name (fqdn). Blank lines
        and lines beginning with "#" are ignored, as is anything on a line
        following a "#".

    -R, --restartServices:
        If all zone requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi authzone add t1.authzone.rice.edu
        Create the "t1.authzone.rice.edu" authoritative zone.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["authzone:delete"] = `
NAME
    ibapi authzone delete - delete Infoblox authoritative zones

USAGE
    ibapi authzone delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox authoritative zones. To
    delete a single authoritative zone, a single zone name (fqdn) may be
    provided as a command line argument. Alternatively, a list of zones to
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
    -V <view>, --View=<view>:
        Specify the view of the zone to delete. Default: "default".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of authoritative zones to
        delete. Each line should contain a hostname and optionally a TXT
        value, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

    -R, --restartServices:
        If all zone requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

    --assumeYes:
        Do not ask for deletion confirmation (assume "yes"). Not
        recommended, except in test environments.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi authzone delete t1.authzone.rice.edu
        Delete the authoritative zone named t1.authzone.rice.edu.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["authzone:get"] = `
NAME
    ibapi authzone get - get Infoblox authoritative zones

USAGE
    ibapi authzone get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox authoritative zones.

    By default, the zone name of each requested zone is shown. The --verbose
    option can be specified to print out the raw API response, which
    includes additional fields.

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
        Specify the view of the zone to fetch. Specify "any" to search for
        zones in all views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the zone(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of authoritative zones to get.
        Each line should contain a hostname and/or a TXT value, separated by
        one or more spaces. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

    -r <obj_ref>, --Ref=<obj_ref>:
        Instead of showing the name and content of the fetched zone(s), show
        the object reference of each zone.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi authzone get t1.authzone.rice.edu
        Fetch the authoritative zone named t1.authzone.rice.edu.

    ibapi authzone get -F fqdn~=authzone.rice.edu -V external
        Fetch the authoritative zone in the external DNS view whose name
        matches authzone.rice.edu.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["authzone:update"] = `
NAME
    ibapi authzone update - update Infoblox authoritative zones

USAGE
    ibapi authzone update <options/args>

DESCRIPTION
    The update command is used to update Infoblox authoritative zones. To
    update a single authoritative zone, a single zone name may be provided
    as a command line argument. Alternatively, a list of zones to update can
    be specified in a file (see --filename).

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
        Specify the view of the zone to update. Default: "default".

    -n <ns_group>, --nsgroup=<ns_group>:
        Update the name server group of the zone.

    -z <zone_format>, --zoneFormat=<zone_format>:
        Update the zone_format of the zone. Valid values are: FORWARD, IPV4,
        and IPV6.

    -D <true|false>, --Disable=<true|false>:
        Update the zone's disabled status to the specified value. Note this
        is not a boolean flag - the value "true" or "false" must be
        specified.

    -c <comment>, --Comment=<comment>:
        Update the zone's comment.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For instance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of authoritative zones to
        update. Each line must contain a hostname and, depending on the
        specified options, a TXT value. Blank lines and lines beginning with
        "#" are ignored, as is anything on a line following a "#".

    -R, --restartServices:
        If all zone requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi authzone update t1.authzone.rice.edu --nsgroup PurdueExternal
        Update the name server group of the specified authoritative zone.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), ibapi.conf(5)


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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        instance: "comment=RT100931",view=default,ttl=900".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:delete(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    -V <view>, --View=<view>:
        Specify the view of the record to delete. Default: "default".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:get(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:update(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        Specify fields and corresponding values to be updated. For instance:
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), fixedaddress(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["fixedaddress"] = `
NAME
    ibapi fixedaddress - create, read, update and delete Infoblox
    fixedaddress records

USAGE
    ibapi fixedaddress <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi fixedaddress" can be used to add, get, update and delete Infoblox
    fixedaddress records. The basic format is

    * ibapi fixedaddress <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi fixedaddress add -h

EXAMPLES
    ibapi fixedaddress add 168.7.56.224 c8:1f:66:c3:a7:e1 -n sitedisk -b
    "/grub2/grubx64.efi" -R
        Add a fixedaddress record with IPv4 168.7.56.224, MAC
        c8:1f:66:c3:a7:e1, name sitedisk, and bootfile "/grub2/grubx64.efi".
        When done, issue the "restart_if_needed" command to restart Grid
        services if needed.

    ibapi fixedaddress delete 168.7.56.224
        Delete a fixedaddress record.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress:add(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["fixedaddress:add"] = `
NAME
    ibapi fixedaddress add - create Infoblox fixedaddress records

USAGE
    ibapi fixedaddress add <options/args>

DESCRIPTION
    The add command is used to create Infoblox fixedaddress records. To
    create a single fixedaddress record, a single IPv4 address and a MAC
    address can be provided as command line arguments, in either order.
    Alternatively, a list of records to add can be specified in a file (see
    --filename).

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
    -n <name>, --Name=<name>
        Specify the name of the new record.

    -V <network_view>, --View=<network_view>:
        Specify the network_view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default:
        "ibapi:fixedaddress:add".

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        instance: "comment=RT100931",network_view=external".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of fixedaddress records to
        create. Each line should contain an IPv4 address and a MAC address,
        in either order, separated by one or more spaces. Blank lines and
        lines beginning with "#" are ignored, as is anything on a line
        following a "#".

    -R, --restartServices:
        If all record requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

    -b <bootfile>, --bootfile=<bootfile>:
        Specify the bootfile for the specified IP address.

    -N <nextserver>, --nextserver=<nextserver>:
        Specify the nextserver for the specified IP address.

    -B <bootserver>, --bootserver=<bootserver>:
        Specify the bootserver for the specified IP address.

    -I <ipFields>, --ipFields=<ipFields>:

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi fixedaddress add 10.143.195.121 c8:1f:66:c1:79:a1 -n
    zabbix-n1.rice.edu
        Create a new fixedaddress record with IPv4 address "10.143.195.121",
        MAC address "c8:1f:66:c1:79:a1", and name "zabbix-n1.rice.edu".

    ibapi fixedaddress add 10.10.10.214 -d 64:00:6a:8f:cc:4d -N10.128.81.10
    -b/grub2/grubx64.efi -R
        Create a fixedaddress record with the specified IPv4 address, MAC
        address, nextserver and bootfile. When done, issue the
        "restart_if_needed" command to restart Grid services if needed.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:delete(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["fixedaddress:delete"] = `
NAME
    ibapi fixedaddress delete - delete Infoblox fixedaddress records

USAGE
    ibapi fixedaddress delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox fixedaddress records. To
    delete a single fixedaddress record, a single IPv4 address and/or MAC
    address may be provided as command line arguments. Alternatively, a list
    of records to delete can be specified in a file (see --filename).

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
        Specify the view of the record to delete. Default: "default".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of fixedaddress records to
        delete. Each line should contain an IPv4 address and/or a MAC
        address, in either order, separated by one or more spaces. Blank
        lines and lines beginning with "#" are ignored, as is anything on a
        line following a "#".

    -R, --restartServices:
        If all record requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi fixedaddress delete 168.7.56.224 c8:1f:66:c1:79:a1
        Delete the fixedaddress record with IPv4 address "168.7.56.224" and
        MAC address "c8:1f:66:c1:79:a1".

    ibapi fixedaddress delete 168.7.56.224
        Delete the fixedaddress record with IPv4 address "168.7.56.224".

    ibapi fixedaddress delete c8:1f:66:c1:79:a1
        Delete the fixedaddress record with MAC address "c8:1f:66:c1:79:a1".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:get(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["fixedaddress:get"] = `
NAME
    ibapi fixedaddress get - get Infoblox fixedaddress records

USAGE
    ibapi fixedaddress get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox fixedaddress records.

    By default, the IP address and MAC address of each fetched record is
    shown. The --verbose option can be specified to print out the raw
    response from the API.

    To fetch a single fixedaddress record, a single IPv4 address and/or MAC
    address may be provided as command line arguments. Alternatively, a list
    of records to fetch can be specified in a file (see --filename).

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
    -V <network_view>, --View=<network_view>:
        Specify the network_view of the record to fetch. Specify "any" to
        search for records in all network_views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of fixedaddress records to get.
        Each line should contain an IPv4 address and/or a MAC address, in
        either order, separated by one or more spaces. Blank lines and lines
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi fixedaddress get -F mac~=00:10:10:10:10 -V external
        Fetch each fixedaddress record (in the external network_view) whose
        MAC address matches "00:10:10:10:10".

    ibapi fixedaddress get -F ipv4addr~=10.143.195.
        Fetch each fixedaddress record (in any network_view) that has an
        IPv4 address matching "10.143.195.".

    ibapi fixedaddress get 00:10:10:10:10:01
        Fetch the fixedaddress record with MAC address "00:10:10:10:10:01".

    ibapi fixedaddress get 168.7.56.224
        Fetch the fixedaddress record with IPv4 address "168.7.56.224".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:update(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["fixedaddress:update"] = `
NAME
    ibapi fixedaddress update - update Infoblox fixedaddress records

USAGE
    ibapi fixedaddress update <options/args>

DESCRIPTION
    The update command is used to update Infoblox fixedaddress records.

    To delete a single fixedaddress record, a single IPv4 address and/or a
    MAC address can be provided as command line arguments, in either order.
    Alternatively, a list of records to delete can be specified in a file
    (see --filename).

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
    -V <network_view>, --View=<network_view>:
        Specify the network_view of the record to update. Default:
        "default".

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -i <new_IPv4_address>, --ip=<new_IPv4_address>:
        Update the IPv4 address of the specified record to
        "new_IPv4_address".

    -m <new_MAC>, --mac=<new_MAC>:
        Update the MAC address of the specified record to "new_MAC".

    -b <bootfile>, --bootfile=<bootfile>:
        Update the bootfile of the specified IP address.

    -N <nextserver>, --nextserver=<nextserver>:
        Update the nextserver of the specified IP address.

    -B <bootserver>, --bootserver=<bootserver>:
        Update the bootserver of the specified IP address.

    -n <new_name>, --Name=<new_name>:
        Update the name of the specified record to "new_name".

    -c <comment>, --Comment=<comment>:
        Update the record's comment.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For instance:
        "comment=RT100931",network_view=default,name=rambo".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of fixedaddress records to
        update. Each line should contain an IPv4 address and a MAC address,
        in either order, separated by one or more spaces. Blank lines and
        lines beginning with "#" are ignored, as is anything on a line
        following a "#".

    -R, --restartServices:
        If all record requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi fixedaddress update 168.7.56.224 -b "/grub2/grubx64.efi"
        Update the bootfile of the fixedaddress record with IPv4
        168.7.56.224.

    ibapi fixedaddress update 168.7.56.224 -i 168.7.56.225
        Update the "168.7.56.224" fixedaddress record, changing the IP
        address to "168.7.56.225".

    ibapi fixedaddress update 168.7.56.224 -m c8:1f:66:c1:79:a1
        Update the MAC address of the "168.7.56.224" fixedaddress record.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), grid(1), host(1), host:add(1), host:delete(1),
    host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["grid"] = `
NAME
    ibapi grid - manage the Infoblox grid

USAGE
    ibapi grid [<options>] <retart|ref>

DESCRIPTION
    "ibapi grid" can be used to instruct Infoblox to restart any grid
    services that need to be restarted, generally due to pending updates
    that require a particular service, such as DHCP, is be restarted.

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
  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi grid restart
        Instruct Infoblox to restart any grid services that need to be
        restarted, generally due to pending updates that require a
        particular service, such as DHCP, is be restarted.

    ibapi grid ref
        Fetch the reference ID of the Infoblox grid.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), host(1), host:add(1),
    host:delete(1), host:get(1), host:update(1), ibapi(1), ptr(1),
    ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    ibapi host add zabbix.rice.edu 168.7.56.225
        Create the "zabbix.rice.edu" Host record with IPv4 address
        "168.7.56.225".

    ibapi host add zabbix.rice.edu 168.7.56.225 -d -R -m f4:8e:38:84:89:e6
    -N 10.128.95.14 -b "/grub2/grubx64.efi"
        Create the "zabbix.rice.edu" Host record with the specified IPv4
        address enabled for DHCP, and when done, issue the
        "restart_if_needed" command to restart Grid services if needed.

    ibapi host update zabbix.rice.edu 168.7.56.225 -i 168.7.56.211
        Update the zabbix.rice.edu IPv4 address "168.7.56.225" to
        "168.7.56.211".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host:add(1),
    host:delete(1), host:get(1), host:update(1), ibapi(1), ptr(1),
    ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        instance: "comment=RT100931",view=default,ttl=900".

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

    -R, --restartServices:
        If all record requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi host add zabbix.rice.edu 168.7.56.225
        Create the "zabbix.rice.edu" Host record with IPv4 address
        "168.7.56.225".

    ibapi host add zabbix.rice.edu 168.7.56.225 -d -R -m f4:8e:38:84:89:e6
    -N 10.128.95.14 -b "/grub2/grubx64.efi"
        Create the "zabbix.rice.edu" Host record with the specified IPv4
        address enabled for DHCP, and when done, issue the
        "restart_if_needed" command to restart Grid services if needed.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:delete(1), host:get(1), host:update(1), ibapi(1), ptr(1),
    ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    -V <view>, --View=<view>:
        Specify the view of the record to delete. Default: "default".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of Host records to delete. Each
        line should contain a hostname and optionally an IP address, in
        either order, separated by one or more spaces. Blank lines and lines
        beginning with "#" are ignored, as is anything on a line following a
        "#".

    -R, --restartServices:
        If all record requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:get(1), host:update(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["host:get"] = `
NAME
    ibapi host get - get Infoblox Host records

USAGE
    ibapi host get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox Host records.

    By default, the hostname and IP address of each fetched record is shown.
    The --verbose option can be specified to print out the raw API response,
    which includes additional fields.

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi host get -F name~=cs.rice.edu -V external
        Fetch each host record (in the external DNS view) whose name matches
        cs.rice.edu.

    ibapi host get -F ipv4addr~=10.143.195.
        Fetch each host record (in any DNS view) that has an IP address
        matching "10.143.195.".

    ibapi host get rb4.rice.edu 168.7.56.224
        Fetch the Host record with hostname "rb4.rice.edu" and IP address
        "168.7.56.224".

    ibapi host get 168.7.56.224
        Get all Host records that contain the IP address "168.7.56.224".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:update(1), ibapi(1), ptr(1),
    ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        Specify fields and corresponding values to be updated. For instance:
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

    -R, --restartServices:
        If all record requests are successfully processed, instruct Infoblox
        to restart any grid services that need to be restarted, generally
        due to pending updates that require a particular service, such as
        DHCP, is be restarted.

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi host update rb4.rice.edu 168.7.56.224 -b "/grub2/grubx64.efi"
        Update the bootfile of IP address 168.7.56.224 of host rb4.rice.edu.

    ibapi host update rb4.rice.edu 168.7.56.224 -i 168.7.56.225
        Update the "rb4.rice.edu" Host record, changing the IP address
        "168.7.56.224" to "168.7.56.225".

    ibapi host update rb4.rice.edu -i +168.7.56.225
        Update the "rb4.rice.edu" Host record, adding the IP address
        "168.7.56.225"

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), ibapi(1), ptr(1), ptr:add(1),
    ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["ibapi"] = `
NAME
    ibapi - Infoblox WAPI command line tool

USAGE
    ibapi <OBJECT> <OPERATION> [<args>]

    where OBJECT is one of

      a alias authzone cname fixedaddress grid host mx txt ptr url

    and OPERATION, for all but the "grid" object, is one of

        add get update delete

DESCRIPTION
    ibapi is a command for adding, reading, updating and deleting basic DNS
    records as well as for managing other Infoblox-specific objects via the
    Infoblox WAPI. Currently supported object types are the DNS records A,
    Alias, CNAME, MX, TXT and PTR; the Infoblox-specific object types
    zoneauth, fixedaddress, grid and host; and the special type "url", which
    allows you to manipulate any type of Infoblox object.

    Use the ibapi -h/--help option for more details. For instance:

    * ibapi -h

    * ibapi host -h

    * ibapi host add -h

Configuration Files
    ibapi configuration files can be used to set defaults for most of the
    available options. The ibapi command searches for configuration files in
    several places, including:

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
    ibapi currently supports only basic authentication. A username and
    password can be specified via command line options or via a
    "username:password" string stored in a local file. Here is an example of
    configuring a username and password for WAPI user "sandman":

    *   mkdir -p $HOME/.ibapi/private

    *   chmod 700 $HOME/.ibapi/private

    *   echo "APIAuthTokenID = sandman" > $HOME/.ibapi/ibapi.conf

    *   echo "sandman:WAPI_PASSWORD" > $HOME/.ibapi/private/sandman

    *   chmod 600 $HOME/.ibapi/private/sandman

Building/Installing
    ibapi is written in Go and is available at github.com. Below is an
    installation example for Fedora/RHEL:

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

EXAMPLES
    ibapi host add zabbix.rice.edu 168.7.56.225 -d -R -m f4:8e:38:84:89:e6
    -N 10.128.95.14 -b "/grub2/grubx64.efi"
        Create the "zabbix.rice.edu" Host record with the specified IPv4
        address enabled for DHCP, and when done, issue the
        "restart_if_needed" command to restart Grid services if needed.

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

    ibapi mx add mail.rice.edu mx1.mail.rice.edu -p 30
        Create an MX record with a preference value of 30.

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

    ibapi grid restart
        Instruct Infoblox to restart any grid services that need to be
        restarted, generally due to pending updates that require a
        particular service, such as DHCP, is be restarted.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ptr(1),
    ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["mx"] = `
NAME
    ibapi mx - create, read, update and delete Infoblox MX records

USAGE
    ibapi mx <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi mx" can be used to add, get, update and delete Infoblox MX
    records. The basic format is

    * ibapi mx <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi mx add -h

EXAMPLES
    ibapi mx add rb4.rice.edu mx1.mail.rice.edu -p 30
        Add an MX record with a preference of 30

    ibapi mx delete rb4.rice.edu mx1.mail.rice.edu
        Delete an MX record.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["mx:add"] = `
NAME
    ibapi mx add - create Infoblox MX records

USAGE
    ibapi mx add <options/args>

DESCRIPTION
    The add command is used to create Infoblox MX records. To create a
    single MX record, a single domain and MX can be provided as command line
    arguments. Alternatively, a list of records to add can be specified in a
    file (see --filename).

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
    -p <preference>, --preference=<preference>:
        Specify the preference for the new record. Default: "10".

    -V <view>, --View=<view>:
        Specify the view for the new record. Default: "default".

    -D, --Disable:
        Disable the new record. Default: false.

    -c <comment>, --Comment=<comment>:
        Specify the comment for the new record. Alternatively, you can
        specify this via the --fields option. Default: "ibapi:mx:add".

    --TTL=<ttl>:
        Specify the ttl for the new record. Alternatively, you can specify
        this via the --fields option.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        instance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of MX records to create. Each
        line should contain a domain and an MX, separated by one or more
        spaces. Blank lines and lines beginning with "#" are ignored, as is
        anything on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi mx add rb4.rice.edu mx1.mail.rice.edu
        Create a new MX record with domain "rb4.rice.edu", MX
        "mx1.mail.rice.edu", and the default preference.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["mx:delete"] = `
NAME
    ibapi mx delete - delete Infoblox MX records

USAGE
    ibapi mx delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox MX records. To delete a
    single MX record, a single hostname and optionally an MX value may be
    provided as command line arguments. Alternatively, a list of records to
    delete can be specified in a file (see --filename).

    If an MX value is specified, the MX record to delete must contain that
    MX value, else no MX record will be deleted. If no MX value is specified
    and only one MX record is found for the specified name, that MX record
    is deleted regardless of its MX value. If multiple MX records are found
    for the same name, the deletion process is aborted (no records are
    deleted) unless the --multiple options is specified to allow mutliple
    record deletions per request.

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
    -p <preference>, --preference=<preference>:
        Specify the preference of the record to delete; only needed when
        multiple MX records exist with the same name and MX value.

    -V <view>, --View=<view>:
        Specify the view of the record to delete. Default: "default".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of MX records to delete. Each
        line should contain a domain and an MX value, separated by one or
        more spaces. Blank lines and lines beginning with "#" are ignored,
        as is anything on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi mx delete rb4.rice.edu mx1.mail.rice.edu
        Delete the MX record with domain "rb4.rice.edu" and MX
        "mx1.mail.rice.edu".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:get(1), mx:update(1), txt(1), txt:add(1), txt:delete(1), txt:get(1),
    txt:update(1), authzone(1), authzone:add(1), authzone:delete(1),
    authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["mx:get"] = `
NAME
    ibapi mx get - get Infoblox MX records

USAGE
    ibapi mx get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox MX records.

    By default, the domain, MX and preference of each fetched record is
    shown. The --verbose option can be specified to print out the raw
    response from the API.

    To fetch MX records, a single domain, or a domain and MX value, may be
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
    -m <mx>, --mx=<mx>:
        Fetch all records with the specified MX value.

    -p <preference>, --preference=<preference>:
        Fetch only records with the specified preference. By default, all MX
        record are shown..

    -V <view>, --View=<view>:
        Specify the view of the record to fetch. Specify "any" to search for
        records in all views. Default: "any".

    -F <fields>, --Fields=<fields>:
        Specify a comma-separated list of field name/value pairs to restrict
        the record(s) fetched.

    -R <return_fields>, --rFields=<return_fields>:
        Specify additional fields to show when in Verbose mode.

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of MX records to get. Each line
        should contain a domain to be deleted. Blank lines and lines
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi mx get rb4.rice.edu -V any
        Fetch the MX records for domain "rb4.rice.edu".

    ibapi mx get -m mx1.mail.rice.edu -V any
        Get the MX records that specify the MX "mx1.mail.rice.edu".

    ibapi mx get rb4.rice.edu mx1.mail.rice.edu -V any
        Fetch the MX records with domain "rb4.rice.edu" and MX
        "mx1.mail.rice.edu".

    ibapi mx get -V any -F mail_exchanger~=mail.rice.edu
        Fetch the MX records in which the MX value matches "mail.rice.edu".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:update(1), txt(1), txt:add(1), txt:delete(1),
    txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["mx:update"] = `
NAME
    ibapi mx update - update Infoblox MX records

USAGE
    ibapi mx update <options/args>

DESCRIPTION
    The update command is used to update Infoblox MX records. To update a
    single MX record, a single hostname and optionally an MX value may be
    provided as command line arguments. Alternatively, a list of records to
    update can be specified in a file (see --filename).

    If an MX value is specified, the MX record to update must contain that
    MX value, else no MX record will be updated. If no MX value is specified
    and only one MX record is found for the specified name, that MX record
    is updated regardless of its MX value. If multiple MX records are found
    for the same request, the update process is aborted (no records are
    updated) unless the --multiple options is specified to allow mutliple
    record updates per request.

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
    -P <preference>, --preference=<preference>:
        Specify the preference of the record to update; only needed when
        multiple MX records exist with the same name and MX value.

    -V <view>, --View=<view>:
        Specify the view of the record to update. Default: "default".

    -n <newName>, --name=<newName>:
        Update the name/domainname value of the specified MX record to
        "newName"

    -m <newMX>, --MX=<newMX>:
        Update the MX value of the specified MX record to "newMX"

    -p <newPreference>, --preference=<newPreference>:
        Update the MX value of the specified MX record to "newPreference"

    -D <true|false>, --Disable=<true|false>:
        Update the record's disabled status to the specified value. Note
        this is not a boolean flag - the value "true" or "false" must be
        specified.

    -c <comment>, --Comment=<comment>:
        Update the record's comment.

    --TTL=<ttl>:
        Update the the record's TTL.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For instance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of MX records to update. Blank
        lines and lines beginning with "#" are ignored, as is anything on a
        line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi mx update rb4.rice.edu mx1.mail.rice.edu -m mx2.mail.rice.edu
        Update the "rb4.rice.edu/mx1.mail.rice.edu" MX record, changing the
        MX value to "mx2.mail.rice.edu".

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), txt(1), txt:add(1), txt:delete(1), txt:get(1),
    txt:update(1), authzone(1), authzone:add(1), authzone:delete(1),
    authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        instance: "comment=RT100931",view=default,ttl=900".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    -V <view>, --View=<view>:
        Specify the view of the record to delete. Default: "default".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:get(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:update(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
        Specify fields and corresponding values to be updated. For instance:
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), url(1), url:add(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["txt"] = `
NAME
    ibapi txt - create, read, update and delete Infoblox TXT records

USAGE
    ibapi txt <add|get|update|delete> <options/args>

DESCRIPTION
    "ibapi txt" can be used to add, get, update and delete Infoblox TXT
    records. The basic format is

    * ibapi txt <operation> <options/args>

    For more details, invoke the specific operation with the --help|-h
    option. For example:

    * ibapi txt add -h

EXAMPLES
    ibapi txt add t1.txt.rice.edu "v=spf1 a:mh.rice.edu
    a:a16.spf.rice.edu/16 -all"
        Add a TXT record.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt:add(1), txt:delete(1),
    txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["txt:add"] = `
NAME
    ibapi txt add - create Infoblox TXT records

USAGE
    ibapi txt add <options/args>

DESCRIPTION
    The add command is used to create Infoblox TXT records. To create a
    single TXT record, a single hostname and TXT value can be provided as
    command line arguments. Alternatively, a list of records to add can be
    specified in a file (see --filename).

    A TXT value can contain spaces and other strange characters that are not
    so URL friendly. ibapi will add double quotes around the TXT data if the
    data is not already quoted. This seems to work well with infoblox, and
    DNS clients get the correct (non-quoted) value. ibapi also escapes
    non-URL friendly characters.

    The maximum size of a string in a TXT record is 255 chars, but a record
    can have multiple strings (the client joins the strings back into one).
    In the Infloblox GUI, this could be done by splitting a long string into
    sub-strings with double quotes: "bigString1" "bigString2". ibapi does
    the same.

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

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values for the new record. For
        instance: "comment=RT100931",view=default,ttl=900".

    -f <filename>, --Filename=<filename>:
        Specify a filename containing a list of TXT records to create. Each
        line should contain a hostname and a TXT value, separated by one or
        more spaces. Blank lines and lines beginning with "#" are ignored,
        as is anything on a line following a "#".

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi txt add t1.txt.rice.edu "v=spf1 a:mh.rice.edu
    a:a16.spf.rice.edu/16 -all"
        Create the "t1.txt.rice.edu" TXT record as specified.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:delete(1),
    txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["txt:delete"] = `
NAME
    ibapi txt delete - delete Infoblox TXT records

USAGE
    ibapi txt delete <options/args>

DESCRIPTION
    The delete command is used to delete Infoblox TXT records. To delete a
    single TXT record, a single hostname and optionally a TXT value may be
    provided as command line arguments. Alternatively, a list of records to
    delete can be specified in a file (see --filename).

    If a TXT value is specified, the TXT record to delete must contain that
    TXT value, else no TXT record will be deleted. If no TXT value is
    specified and only one TXT record is found for the specified name, that
    TXT record is deleted regardless of its TXT value. If multiple TXT
    records are found for the same request, the deletion process is aborted
    (no records are deleted) unless the --multiple options is specified to
    allow mutliple record deletions per request.

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
        Specify the view of the record to delete. Default: "default".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of TXT records to delete. Each
        line should contain a hostname and optionally a TXT value, separated
        by one or more spaces. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

    -m, --multiple:
        If only a name is specified (no TXT value is specified), allow
        deletion of multiple records if multiple records are found for the
        specified name. This option has no effect if both the name and data
        value are specified.

  OPTIONS - API Options
    --APIBaseURL=<url>:
        API base URL. Default: "https://infoblox.rice.edu/wapi/v2.11"

    --HTTPTimeout=<seconds>:
        Timeout in seconds for the HTTP connection. Default: 10.

    --APIAuthMethod=<method>:
        WAPI authentication method for accessing the Infoblox API.
        Currently, only "Basic" authentication (username and password) is
        supported.

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi txt delete t1.txt.rice.edu "v=spf1 a:mh.rice.edu
    a:a16.spf.rice.edu/16 -all"
        Delete the TXT record with hostname "t1.txt.rice.edu" and TXT value
        "v=spf1 a:mh.rice.edu a:a16.spf.rice.edu/16 -all".

    ibapi txt delete t1.txt.rice.edu
        Delete all TXT records named t1.txt.rice.edu.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1), txt:get(1),
    txt:update(1), authzone(1), authzone:add(1), authzone:delete(1),
    authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["txt:get"] = `
NAME
    ibapi txt get - get Infoblox TXT records

USAGE
    ibapi txt get <options/args>

DESCRIPTION
    The get command is used to read/fetch Infoblox TXT records.

    By default, the hostname and TXT value of each fetched record is shown.
    The --verbose option can be specified to print out the raw API response,
    which includes additional fields.

    To fetch a single TXT record, a single hostname and/or TXT value may be
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
        Specify a filename containing a list of TXT records to get. Each
        line should contain a hostname and/or a TXT value, separated by one
        or more spaces. Blank lines and lines beginning with "#" are
        ignored, as is anything on a line following a "#".

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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi txt get t1.txt.rice.edu
        Fetch each TXT record named t1.txt.rice.edu.

    ibapi txt get -F name~=txt.rice.edu -V external
        Fetch each TXT record (in the external DNS view) whose name matches
        t1.txt.rice.edu.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


`
	PodMap["txt:update"] = `
NAME
    ibapi txt update - update Infoblox TXT records

USAGE
    ibapi txt update <options/args>

DESCRIPTION
    The update command is used to update Infoblox TXT records. To update a
    single TXT record, a single hostname and optionally a TXT value may be
    provided as command line arguments. Alternatively, a list of records to
    update can be specified in a file (see --filename).

    If a TXT value is specified, the TXT record to update must contain that
    TXT value, else no TXT record will be updated. If no TXT value is
    specified and only one TXT record is found for the specified name, that
    TXT record is updated regardless of its TXT value. If multiple TXT
    records are found for the same request, the update process is aborted
    (no records are updated) unless the --multiple options is specified to
    allow mutliple record updates per request.

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

    -t <TXT>, --txt=<TXT>:
        Update the record's TXT value.

    -F <fields>, --Fields=<fields>:
        Specify fields and corresponding values to be updated. For instance:
        "comment=RT100931",view=default,ttl=900".

    -f <filename>, --filename=<filename>:
        Specify a filename containing a list of TXT records to update. Each
        line must contain a hostname and, depending on the specified
        options, a TXT value. Blank lines and lines beginning with "#" are
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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    ibapi txt update t1.txt.rice.edu "v=spf1 a:mh.rice.edu
    a:a16.spf.rice.edu/16 -all" -t "v=spf1 a:mh.rice.edu
    a:a16.spf.rice.edu/16 -all"
        Update the TXT value of the specified TXT record.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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
    ibapi url get 'zone_auth?fqdn~=mail.rice.edu'
        Retrieve all "authoritative" zones that match "mail.rice.edu".

    ibapi url get 'nsgroup?name~=.'
        Retrieve all name server groups.

    ibapi url get '/record:host?name~=cs.rice.edu'
        Retrieve all Host records with a name that matches the pattern
        "cs.rice.edu".

    ibapi url get '/record:a?ipv4addr~=128.42.201.'
        Retrieve all A records with an IP address that matches the pattern
        "128.42.201.".

    ibapi url get '/record:host?_schema'
        Retrieve the schema for a Host record.

FILES
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1),
    url:add(1), url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:delete(1), url:get(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:get(1), url:update(1), mx(1), mx:add(1), mx:delete(1),
    mx:get(1), mx:update(1), txt(1), txt:add(1), txt:delete(1), txt:get(1),
    txt:update(1), authzone(1), authzone:add(1), authzone:delete(1),
    authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:update(1), mx(1), mx:add(1),
    mx:delete(1), mx:get(1), mx:update(1), txt(1), txt:add(1),
    txt:delete(1), txt:get(1), txt:update(1), authzone(1), authzone:add(1),
    authzone:delete(1), authzone:get(1), authzone:update(1), ibapi.conf(5)


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

    --username=<username>:
        Specify the username used for basic auth.

    --password=<password>:
        Specify the password used for basic auth. If this option is
        specified and is non-empty, either the --username option can be used
        to specify the corresponding username, or the current user will be
        assumed for username.

    --PromptForPassword:
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

    --GridReference=<grid_reference>:
        Specify the Infoblox grid reference ID. This can be used to save a
        fetch when the --restartServices option is specified. While this
        option is only relevant to a few commands, it is allowed (ignored)
        by the other commands.

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
    /usr/site/ibapi-1.0/etc/ibapi.conf
    /etc/opt/ibapi/ibapi.conf
    /etc/opt/ibapi-1.0//ibapi.conf
    ~/.ibapi/ibapi.conf
    ~/.ibapi-1.0/ibapi.conf
        The IBAPI configuration files which can be used to set defaults for
        nearly all of the options described above. Any combination of these
        may be used. Each file found is read in turn, with settings in later
        files overriding those in previous files. Note that command line
        options override all config file settings.

SEE ALSO
    a(1), a:add(1), a:delete(1), a:get(1), a:update(1), alias(1),
    alias:add(1), alias:delete(1), alias:get(1), alias:update(1), cname(1),
    cname:add(1), cname:delete(1), cname:get(1), cname:update(1),
    fixedaddress(1), fixedaddress:add(1), fixedaddress:delete(1),
    fixedaddress:get(1), fixedaddress:update(1), grid(1), host(1),
    host:add(1), host:delete(1), host:get(1), host:update(1), ibapi(1),
    ptr(1), ptr:add(1), ptr:delete(1), ptr:get(1), ptr:update(1), url(1),
    url:add(1), url:delete(1), url:get(1), mx(1), mx:add(1), mx:delete(1),
    mx:get(1), mx:update(1), txt(1), txt:add(1), txt:delete(1), txt:get(1),
    txt:update(1), authzone(1), authzone:add(1), authzone:delete(1),
    authzone:get(1), authzone:update(1), ibapi.conf(5)


`
return nil
}

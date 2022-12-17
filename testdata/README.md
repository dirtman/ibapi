## Command line testing

The Go package "testscript" is used for testing.  The tests are invoked via
test_test.go, and the test scripts are in testdata/script/*.  The tests rely on
access to the Infoblox WAPI and permissions to configure records in
"seci.rice.edu".  To test all scripts in testdata/script/:

  go test

To test individual scripts:

  go test -run Test/a
  go test -run Test/host
  go test -run Test/ptr
  go test -run Test/cname
  go test -run Test/alias
  go test -run Test/mix

Double quotes can cause issues with tests.  For instance, for the test

   ibapi .....  -c "Default View"

the ibapi command get 3 args:

  -c
  "Default
  View"

Use single quotes instead.


## Check if any of our "test" records already exist (should get all NOTFOUNDs).

### a1-a5;  201-209
ibapi a get a1.seci.rice.edu 
ibapi a get a2.seci.rice.edu 

### host1-host5;  211-219
ibapi host get host1.seci.rice.edu 
ibapi host get host2.seci.rice.edu 
ibapi host get host3.seci.rice.edu 
ibapi host get host4.seci.rice.edu 
ibapi host get host5.seci.rice.edu 

ibapi cname get cname1.seci.rice.edu
ibapi cname get cname2.seci.rice.edu
ibapi cname get alias1.seci.rice.edu

ibapi alias get alias1.seci.rice.edu
ibapi alias get alias1.seci.rice.edu -TTXT
ibapi alias get alias2.seci.rice.edu -TTXT
ibapi alias get cname1.seci.rice.edu

### 221-129
ibapi ptr get ptr1.seci.rice.edu
ibapi ptr get ptr2.seci.rice.edu
ibapi ptr get ptr3.seci.rice.edu

### any 231-239
ibapi host get host-a.seci.rice.edu
ibapi host get host-b.seci.rice.edu
ibapi a get host-a.seci.rice.edu 10.10.10.231
ibapi a get host-a.seci.rice.edu 10.10.10.231
ibapi a get a-a.seci.rice.edu 10.10.10.231
ibapi ptr get host1.seci.rice.edu 10.10.10.231
ibapi ptr get a-a.seci.rice.edu 10.10.10.231
ibapi ptr get ptr-a.seci.rice.edu 10.10.10.231
ibapi cname get host-a.seci.rice.edu -V external
ibapi cname get cname-a.seci.rice.edu
ibapi alias get alias-a.seci.rice.edu
ibapi alias get alias-txt.seci.rice.edu -T TXT

### url 241-249
ibapi a get url1.seci.rice.edu 
ibapi a get url2.seci.rice.edu 

## To delete the test records:

ibapi a delete a1.seci.rice.edu 10.10.10.101 
ibapi a delete a1.seci.rice.edu 10.10.10.101 -V external
ibapi a delete a1.seci.rice.edu 10.10.10.102 -V external
ibapi a delete a2.seci.rice.edu 10.10.10.101 -V external
ibapi a delete a2.seci.rice.edu 10.10.10.102 -V external

ibapi host delete host1.seci.rice.edu
ibapi host delete host1.seci.rice.edu -V external
ibapi host delete host2.seci.rice.edu -V external
ibapi host delete host3.seci.rice.edu 
ibapi host delete host4.seci.rice.edu 
ibapi host delete host5.seci.rice.edu 

ibapi cname delete cname1.seci.rice.edu
ibapi cname delete cname1.seci.rice.edu -V external
ibapi cname delete cname2.seci.rice.edu
ibapi cname delete alias1.seci.rice.edu

ibapi alias delete alias1.seci.rice.edu
ibapi alias delete alias1.seci.rice.edu -TTXT
ibapi alias delete alias1.seci.rice.edu -TTXT -V external
ibapi alias delete alias2.seci.rice.edu -TTXT -V external
ibapi alias delete cname1.seci.rice.edu

ibapi ptr delete a1.seci.rice.edu 10.10.10.101
ibapi ptr delete a1.seci.rice.edu 10.10.10.102
ibapi ptr delete a2.seci.rice.edu 10.10.10.102
ibapi ptr delete a2.seci.rice.edu 10.10.10.101
ibapi ptr delete a3.seci.rice.edu 10.10.10.102

ibapi host delete host-a.seci.rice.edu
ibapi host delete host-b.seci.rice.edu
ibapi a delete host-a.seci.rice.edu 10.10.10.221
ibapi a delete host-a.seci.rice.edu 10.10.10.231
ibapi a delete a-a.seci.rice.edu 10.10.10.231
ibapi ptr delete host1.seci.rice.edu 10.10.10.231
ibapi ptr delete a-a.seci.rice.edu 10.10.10.231
ibapi ptr delete ptr-a.seci.rice.edu 10.10.10.231
ibapi cname delete host-a.seci.rice.edu -V external
ibapi cname delete cname-a.seci.rice.edu
ibapi alias delete alias-a.seci.rice.edu
ibapi alias delete alias-txt.seci.rice.edu -T TXT

ibapi a delete url1.seci.rice.edu 10.10.10.241
ibapi a delete url2.seci.rice.edu 10.10.10.242 -V external

-----------------------------------------------------------------------------------

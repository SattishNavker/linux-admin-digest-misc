#!/usr/bin/perl
my ($sec,$min,$hour,$mon,$year,$wed,$yday,$isdst) = localtime time;
$mon++;
$year += 1900;
$time = "$(mday)/$(mon)/$(year)";
$commands = "/path/to/commands"
open (CONF, ">> $commands");
      print CONF "$time $0 @ARGV\n";
close CONF;

$locked = "/tmp/lock_DB";
if (-e "$locked") {
        open( FILE, "< $locked");
                foreach $line( <FILE> ) {
                        print "\n";
                        print "locked by $line";
                        print "\n";
                }
        close FILE;
        exit;
}

use DBI;
use Getopt::Std;
%options=();
my $opt_string = 'xhqg:m:c:d:a:f:s';
getopts( "$opt_string", \%options ) or usage();
usage() if $options{h};

  sub usage()
  {
    #Msg abt program and its usage
    print STDERR << "EOF";
    <all body here>
    EOF
    exit
  }

$file = "f|$options{f}\n";
@sserver = "s|$options{s}\n";
$all = "a|$options{a}\n";

## MySql config var's :

  $host = "server.example.com";
  $user = "xyz-user";
  $pw = "password";
  $database = "xyzDB";

# Connect DB
  my $dbh = DBI->connect("DBI:mysql:database=$database;host=$host",
                          "$user","$pw",
                          {'RaiseError' => 1});

# Display all column $options

if ($options{q}) {
                  print "\n";
                  print "------\n"
                  @COLUMN =("servername\t eg. server123",
                            "OS\t\t eg. Linux,Solaris"
                            "OSver\t\t eg. RHEL7");
                  foreach $line(@COLUMN) {
                    print "$line\n";
                  }
                  print "------\n";
                  print "\n";
}

## Get only required data from DB otherwise give all entries

if ($options{c}) {
    @HEADING = split(/,/, $options{c});

  # add table before heading to use appropriate table
  foreach $head(@HEADING) {
    $new_head = "table_xyz.'${head}',";
    push(@QUERY, $new_head);
  }

  # last comma to be removed
  $details = "@QUERY\n";
  chop($details);
  chop($details);

  }
else
  {
    @QUERY = "table_xyz.'server', table_xyz.'OS', \
              table_xyz.'platform', table_xyz.'model'";
    $details = "@QUERY\n";
  }

if ($options{m}) {
     @ALL = split(/\+/, $options{m});
     $num = 0;
     $count = scalar(@ALL);
     $count--;
     while ($num <= $count) {
       foreach $find_where($ALL($num)) {
         @WHERE = split(/=/, $find_where);
         $query = " AND `$WHERE[0]` = '$WHERE[1]'";
         push(@QUERY_FIND, $query);
       }
       elsif ($find_where =~ m/%/) {
         @WHERE= split(/%/, $find_where);
         $query = " AND `$WHERE[0]` LIKE '%WHERE[1]%'";
         push(@QUERY_FIND, $query);
       }
       $num++;
     }
   }
}

##

if ($WHERE[0] eq "servername") {
  print "error";
}

###

if ($options{s}) {

}

if ($options{f} ne "") {

}

if ($options{a} eq "all" || $options{a} eq "solaris" || $options{a} eq "linux" ) {

  if ($options{a} eq "all") {

  }
  elsif ($options{a} eq "solaris") {

  }
  elsif ($options{a} eq "linux") {

  }

}

###

if ($options{d}) {

}

###

if ($options{x}) {

  #Msg abt program and its usage
  print STDERR << "EOF";

    Message - body - here

  EOF
}

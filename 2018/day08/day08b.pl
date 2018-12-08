#!/usr/bin/perl

use warnings;
use strict;

use feature 'say';

use List::Util qw(max min reduce sum);
use List::MoreUtils qw(firstidx firstval pairwise singleton);
use Algorithm::Combinatorics qw(variations);
use Math::Prime::Util qw(is_prime);
use Data::Dumper;
$Data::Dumper::Sortkeys = 1;

sub makenode {
	my ($arr, $idx) = @_;
	my ($childCount, $metaCount) = ($arr->[$$idx], $arr->[$$idx+1]);
	$$idx += 2;
	my $node = {};
	push @{ $node->{children} }, makenode($arr, $idx) while $childCount--;
	push @{ $node->{meta} }, $arr->[$$idx++] while $metaCount--;

	if (not exists $node->{children}) {
		$node->{value} = sum @{ $node->{meta} };
	}
	else {
		foreach my $childIdx (@{ $node->{meta} }) {
			next if $childIdx > @{ $node->{children} };
			$node->{value} += $node->{children}[$childIdx-1]{value};
		}
	}

	return $node;
}

my $fname = shift;

open my $fh, "<", $fname
    or die "Can't open $fname: $!";

my $line = <$fh>;
chomp $line;
my @arr = split / /, $line;

my $idx = 0;

my $tree = makenode(\@arr, \$idx);

say Dumper($tree);
say $tree->{value};

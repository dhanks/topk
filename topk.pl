use strict;
use warnings;

# Returns the k most frequent elements in the given array
sub k_most_frequent {
  my ($nums, $k) = @_;

  # Return an empty array if the input array is empty
  return [] if @$nums == 0;

  # Return the entire array if k is greater than the number of elements
  return @$nums if $k > @$nums;

  # Use a hash to count the frequency of each element
  my %count;
  foreach my $num (@$nums) {
    $count{$num}++;
  }

  # Create a min-heap of elements and their frequencies, sorted by frequency
  my @heap;
  foreach my $num (keys %count) {
    push @heap, [$num, $count{$num}];
  }
  @heap = sort { $a->[1] <=> $b->[1] } @heap;

  # Pop the top k elements off the heap
  my @result;
  while (@result < $k) {
    push @result, shift @heap;
  }

  # Sort the resulting array by frequency in descending order
  @result = sort { $b->[1] <=> $a->[1] } @result;

  # Return the elements
  return map { $_->[0] } @result;
}

# Unit tests
my @tests = (
  # Test empty array
  {
    nums => [],
    k => 5,
    expected => [],
  },
  # Test array with fewer elements than k
  {
    nums => [1, 2, 3],
    k => 5,
    expected => [1, 2, 3],
  },
  # Test array with more elements than k
  {
    nums => [1, 2, 2, 3, 3, 3],
    k => 2,
    expected => [3, 2],
  },
 
);

foreach my $test (@tests) {
  my $result = k_most_frequent($test->{nums}, $test->{k});
  if ($result ~~ $test->{expected}) {
    print "PASS\n";
  } else {
    print "FAIL\n";
  }
}

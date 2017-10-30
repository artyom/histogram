Command histogram reads non-negative float64 numbers on stdin and prints
ASCII histogram of their distribution

Inspired by mmhistogram tool from https://blog.cloudflare.com/three-little-tools-mmsum-mmwatch-mmhistogram/

Example output:

	min:3090 mean:492539.99 median:486079.50 max:989355 stddev:309009.12 cnt:100
	      bkt   --------------------------------------------------   cnt    %
		0                                                          0    0
		1                                                          0    0
		2                                                          0    0
		4                                                          0    0
		8                                                          0    0
	       16                                                          0    0
	       32                                                          0    0
	       64                                                          0    0
	      128                                                          0    0
	      256                                                          0    0
	      512                                                          0    0
	     1024                                                          0    0
	     2048                                                    *     1    1
	     4096                                                    *     1    1
	     8192                                                    *     1    1
	    16384                                                    *     1    1
	    32768                                                    *     1    1
	    65536                                                *****    10   10
	   131072                                              *******    15   15
	   262144                                         ************    25   25
	   524288                               **********************    45   45

Command histogram reads non-negative float64 numbers on stdin and prints
ASCII histogram of their distribution

Inspired by mmhistogram tool from https://blog.cloudflare.com/three-little-tools-mmsum-mmwatch-mmhistogram/

Example output:

	min:3090.00 mean:492539.99 median:486079.50 max:989355.00 stddev:309009.12 cnt:100
	      bkt   --------------------------------------------------   cnt
	     2048                                                    *     1
	     4096                                                    *     1
	     8192                                                    *     1
	    16384                                                    *     1
	    32768                                                    *     1
	    65536                                                *****    10
	   131072                                              *******    15
	   262144                                         ************    25
	   524288                               **********************    45

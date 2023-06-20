## Bitrange

- Transform time ranges to a unsigned 64-bits integer to check two or multi time range are overlap or not

For example, each `-` is 30 minutes block, so we have below block presentation of time in day:

0:00`--------------------------------------------`24:00

Each `-` mean duration from start time until before end time moment (start <= block < end)<br />
Each `1` is `-`, but mark this block is used, `0` is reverted mean.<br />
So if you need present a day with time ranges 8h-12h and 14h-18h, you will have below presentation:

P1: `0000000000000000111111110000111111110000000000000000000000000000`

Now we can store above presentation by an unsigned 64-bits int with value is 280443916124160.
Let's try present other time ranges in day 1h-8h30:

P2: `0011111111111111100000000000000000000000000000000000000000000000` (4611545280939032026)<br />
P1: `0000000000000000111111110000111111110000000000000000000000000000` (280443916124160)

In order to check that [8h-12h and 14h-18h] and [1h-8h30] are overlap or not, we can check each same position bit in P1 and P2<br />
If both bits are `1`, [8h-12h and 14h-18h] and [1h-8h30] are overlap<br />
We can do this easily with bit operation `AND` (&).<br />
You can create P12 = P1 OR P2 then check with P3...

This repository implement some function process above concept in golang. You can try more complex test case on file `bitrange_test.go`. You can change block duration to multiple of 5. 
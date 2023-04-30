</p>
<h2>Not a deaf phone</h2>
It is possible that you once played the game "Deaf Phone", or heard about it. In this game, participants have to transmit information to each other in various ways: verbally, figuratively, sometimes they even have to write a text with their left hand that another team member will have to read. It is also known that the transmitted information almost never reaches the final destination. Denote by Fi(x) a function that converts the text of the transmitted information x into the one that participant i+1 will receive from participant i. Then the last nth participant will receive data y, which will be expressed by the following formula:<br>
<h3>y = Fn-1(Fn-2(…F2(F1(x))))</h3>
But you need to exclude any external factors that can distort the original information and you need to implement a "non-deaf phone" program that can accurately deliver the original data, i.e. in our case, the function Fi(x) = x for all i from 1 to n-1.<br>
<h3>Input data</h3>
In the single line of the input file INPUT.TXT a natural number from 1 to 100 is written.
<h3>Output data</h3>
To the output file OUTPUT.TXT you need to output exactly the same number that is specified in the input file.
</p>
# Prepend | Postpend 

Just a simple tool for concatenating strings in your terminal.

With `prep` the argument will be prepended to each line.

With `posp` the argument will be postpended to each line.

Instead of using stdin you can specify a file with `-i`.

You can also write the concatenation to a file with `-o`.

Use like this:

``
echo "Hello" | prep -s "_" "World" //output: Hello_World
``



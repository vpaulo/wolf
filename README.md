# wolf
Learning how to build a programming language

## Stage 0
Learn the basics of building your own programing language following, How to build an interpreter in go, then optimize for the needs on the scripting language to be built.

## Stage 1
This language aims to be simple and easy to learn, it should be easy to install and update compiler.

### Phase 1
[] Create a new rust package to initialise the Wolf scripting language
[] Create initial documentation showing a draft of the language capabilities
[] Create the initial draft of grammar/syntax of the language
[] Use pest to create the language grammar and tokens
[] Create initial cli to help with development and testing

### Phase 2
[] Build interpreter
[] Build compiler

At the end  of this phase you should have a working programming language with the bare minimun functionality, no standart library.

### Phase 3
[] Build garbage collector
[] Build JIT

we should be able to execute the script in the command line using the JIT

### Phase 4
[] Build AOT

we should be able to create executables of the script

### Phase 5
[] Create standart library for string utilities
[] Create standart library for number utilities
[] Create standart library for array utilities
[] Create standart library for hashmap utilities

## Stage 2
This stage is to create your own "browser" like application that will have an input that will accept a link to a script in our language and draw ui.



Wolf language Draft:
this syntax is in early stages, it will change

// Variable Declaration
var number: number = 1; // number is an alias for i32 or f64
var string: string = "string";

// types can be auto detected
var number = 2;
var string = "hello cecile";


// Array & Builtin function
var arr = [1, 2, 3, 4]; // same as var arr: number[] = [1, 2, 3, 4];
arr.push(5);

println(arr); // Out: [1, 2, 3, 4, 5]

// Hashmap & Builtin function
var map = { 
  "a": 1, 
  "b": 2, 
  "c": 3, 
  "d": 4 
};

// Or
var map: string[number] = { 
  "a": 1, 
  "b": 2, 
  "c": 3, 
  "d": 4 
};

// Function Declaration
fn hello() -> String ::
  return "hello";
end

println(hello()); // Out: "hello"

// Or
var greet = fn() -> String ::
  return "hello";
end

println(greet()); // Out: "hello"

// If Declaration
if str == "hello" ::
  println(str);
end

if str and msg ::
  println(foo);
end

if str or msg ::
  println(foo);
end

if not str ::
  println(foo);
end

// Case Declaration
case str == "hello" ::
  true => println(str);
  false => println("");
end

var msg = case str == "hello" ::
  "hello" => str + " world!";
  _ => str;
end

println(msg) // Out: "hello world!" or "hello"

// For Declaration
var arr = [1, 2, 3, 4];

for value in arr ::
 println(value);
end

for index, value in arr ::
 println(index);
 println(value);
end

var map = {
  "a": 1, 
  "b": 2, 
  "c": 3, 
  "d": 4 
};
for key, value in map ::
 println(key);
 println(value);
end


// Type Declaration
type Area = number;
type Message = string;
type List = string[];
type Items = string[string];

// Struct Declaration
struct Point ::
  x: int,
  y: int,
end

impl Point ::
  fn new(x: int, y: int) ::
    self.x = x;
    self.y = y;
  end
end

// OR

fn Point.new(x: int, y: int) ::
  self.x = x;
  self.y = y;
end

var point: Point = Point();

















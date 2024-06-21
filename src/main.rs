use pest::Parser;
use pest_derive::Parser;
use wolf::cli::Cli;
use std::fs;

#[derive(Parser)]
#[grammar = "./grammar/wolf.pest"]
pub struct WolfParser;

fn main() -> anyhow::Result<()>  {
    let _ = Cli::run();
    let unparsed_file = fs::read_to_string("./example/test.wl").expect("cannot read file");

    let file = WolfParser::parse(Rule::program, &unparsed_file)
        .expect("unsuccessful parse") // unwrap the parse result
        .next()
        .unwrap(); // get and unwrap the `file` rule; never fails

        println!("file: \n{:#?}", file.clone());

    for line in file.clone().into_inner() {
        println!("Pair: \n{:#?}", line);
    }

    Ok(())
}


// TODO look into TEO lang
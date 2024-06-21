use anyhow::Result;
use clap::Parser;
use std::io::{self, Write};

#[derive(Parser)]
#[command(
    name = "wolf",
    version = env!("CARGO_PKG_VERSION"),
    author = env!("CARGO_PKG_AUTHORS"),
    about = "Wolf, a simple programming language"
)]
pub enum Commands {
    Run {
        /// The file path
        #[arg(value_hint = clap::ValueHint::FilePath, value_name = "INPUT FILE PATH")]
        file: String,
        /// Runtime stack trace
        #[arg(short, long)]
        trace: bool,
        /// Opcodes debug
        #[arg(short, long)]
        debug: bool,
        /// AST debug
        #[arg(short, long)]
        ast_debug: bool,
        /// Optimized mode
        #[arg(short, long)]
        optimized: bool,
    }
}

impl Commands {
    pub fn run(&self) -> Result<()> {
        match self {
            Commands::Run {
                file,
                trace,
                debug,
                ast_debug,
                optimized,
            } => {
                let source = std::fs::read_to_string(file)?;

                println!("source: \n{:#?}", source);

                // let mut color = termcolor::StandardStream::stderr(termcolor::ColorChoice::Always);

                // let stdout = &mut io::stdout().lock();

                // let mut allocation = CeAllocation::new();
                // let mut compiler = Compiler::new(&mut allocation, *debug);
                // let mut vm = crate::vm::VM::new(&mut allocation, *trace, *optimized);
                // if let Err(e) = vm.run(&source, &mut color, *ast_debug, stdout, &mut compiler) {
                //     report_err(&source, e);
                // }
                Ok(())
            }
        }
    }
}

// fn report_err(source: &str, errors: Vec<ErrorS>) {
//     let mut buffer = termcolor::Buffer::ansi();
//     for err in errors {
//         report_error(&mut buffer, source, &err);
//     }
//     io::stderr()
//         .write_all(buffer.as_slice())
//         .expect("failed to write to stderr");
// }
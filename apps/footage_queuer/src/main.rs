fn main() {
    // pull in cmdline args
    let args: Vec<String> = std::env::args().collect();
    dbg!(args);

    // get QUEUE_NAME env var
    let queue_name = std::env::var("QUEUE_NAME");
    dbg!(queue_name);
}
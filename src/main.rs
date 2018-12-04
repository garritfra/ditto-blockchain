mod core;

fn main() {
    let mut blockchain = core::blockchain::new();
    let block = core::blockchain::Block { data: 42 };
    blockchain.add_block(block);

    for block in blockchain.get_blocks() {
        println!("{}", block.data);
    }
}

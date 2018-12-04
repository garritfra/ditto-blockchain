mod core;

fn main() {
    let mut blockchain = core::blockchain::new();
    let block = core::blockchain::Block { data: 42 };
    blockchain.addBlock(block);

    for block in blockchain.getBlocks() {
        println!("{}", block.data);
    }
}

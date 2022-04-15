script {    
    use 0x2::Meta42;    
    
    fun meta42_initialize( sig: signer) {
        Meta42::initialize(&sig);        
    }
}

script {    
    use 0x2::Meta42;    
    
    fun meta42_accept( sig: signer) {
        Meta42::accept(&sig);        
    }
}

script {    
    use 0x2::Meta42;    
    
    fun meta42_mint_token( sig: signer, hdfs: vector<u8>) {
        Meta42::mint_token(&sig, hdfs);        
    }
}

script {    
    use 0x2::Meta42;    
    
    fun meta42_share_token_by_id( sig: signer, receiver: address, token_id: vector<u8>, message: vector<u8>) {
        Meta42::share_token_by_id(&sig, receiver, token_id, message );        
    }
}

script {    
    use 0x2::Meta42;    
    
    fun meta42_share_token_by_index( sig: signer, receiver: address, index: u64, message: vector<u8>) {
        Meta42::share_token_by_index(&sig, receiver, index, message );        
    }
}
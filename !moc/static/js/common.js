"use strict";

$(function() {
    $(document).ready(function() {
        var command = [];
        var key_hash = "2767cbd1e98a5d8ea50c49a539bb29c335b3658bc0e9232335ad90a552fff9b6"
        var key = "iamadministrator";
        document.addEventListener('keypress', (event) => {
            var key_name = event.key;
            command.push(key_name);
        
            if(command.length > 16){
                command.shift()
            }
        
            if(get_all_key(command) == key){
                window.open("login.html", "_blank");
            }
        
            function get_all_key(str){
                var key = "";
                for(var i in str){
                    key += str[i];
                }
                return key;
            }
        });
    });
});


import { encode, decode } from "@msgpack/msgpack";

// So Replacing JSON With MessagePack Turns 277 Bytes Into 254 Bytes for the data of this particular endpoint
// Deno runtime (not node.js)

setInterval(function(){
    const s = new Date()
fetch("http://localhost:8080/process-resources/1", {})
.then(res=>res.arrayBuffer())
.then(data=>{
    console.log(
        // decode(data),
        `Fetched & Deserialized in ${new Date().getTime() - s.getTime()}ms`
    )
})
}, 1000)
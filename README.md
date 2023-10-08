# sessint-compiler

Session types compiler - Master's Dissertation

Abstract:
Concurrent programming has grown significantly in popularity thanks to the increased demands for availability and responsiveness placed on software systems to accommodate expanding end-user populations.
Nonetheless, dealing with the inherent challenges of concurrent programming is considerably more complex than in traditional sequential programming.
To gain a better grasp of concurrency, programming languages should incorporate novel functionalities that offer programmers high-level concurrency primitives.

Session types provide a typing discipline to ensure that channel communication follows specific protocols, enforced during type-checking, preventing communication errors.
Additionally, session types based on linear logic, which are foundational to our work, ensure the absence of deadlocks.
Session types aid in analyzing concurrency by representing such systems as a set of processes linked by channels, facilitating process coordination.

As a product of prior research, a compiler has been developed that compiles a functional language utilizing session types into Go.
Our current study introduces an optimization layer to the compiler, which detects opportunities for optimization, enhancing the efficiency of the code generated, focusing on enhancing synchronization properties and improving message exchange. We have implemented three optimizations:

- the forwarding optimization simplifies the redirection of messages between channels, reducing synchronization points and the creation of channels;
- recursion optimization converts recursive functions into loops to minimize the overhead of recursive function calls;
- packing messages optimization combines multiple messages into a single message to decrease the number of messages exchanged and minimize synchronization costs.

After implementing the optimizations and comparing them to the non-optimized version, we conducted a quantitative evaluation to determine performance gains. We also conducted an informal and experimental validation, arguing that the optimizing transformations preserve program semantics.

We have concluded that the optimizations incorporated in this research enhance the performance of the generated code, while also maintaining program semantics and not hindering compilation times. The most notable optimization is the recursion optimization, which can be 33 times faster in certain cases.

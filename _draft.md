~@[meta]
#asm #syscall #x86-64 #call #stack

@[title:title]
Function Call — What Actually Happens (x86-64 SysV ABI)

---

@[b001] tldr x86 sysv abi
A function call on x86-64 (SysV ABI) is not just a jump.

It modifies the stack, stores a return address, optionally saves the previous base pointer, and transfers execution to a new instruction pointer. Understanding this is foundational to ABI, stack frames, and debugging.

---

@[b002] def
A function call is a control transfer operation that:

1. Saves the return address.
2. Optionally saves the current base pointer.
3. Adjusts the stack pointer.
4. Transfers execution to the callee’s entry point.

---

~@[e001] example stack state
Before `call add`

Registers:
- RSP = 0x7ffdf000
- RIP = 0x400540

Stack (top grows downward):

0x7ffdf000  ← RSP

After `call add`

CPU performs:

RSP = RSP - 8  
[RSP] = return_address  
RIP = &add  

New state:

- RSP = 0x7ffdeff8  
- [0x7ffdeff8] = 0x400545 (next instruction in main)  
- RIP = address of add

---

@[b003] mechanism
Consider this C code:

```c
int add(int a, int b) {
    return a + b;
}

int main() {
    int x = add(2, 3);
    return 0;
}
```

On x86-64 (SysV ABI), the call roughly translates to:

```asm
mov edi, 2
mov esi, 3
call add
```

The call instruction performs two atomic actions:
	1.	Pushes the return address onto the stack.
	2.	Sets RIP to the address of add.

{[e001]}

---

@[b004:example] mechanism prologue
Typical function prologue:

```asm
push rbp
mov rbp, rsp
sub rsp, 16
```

This:
	1.	Saves the old base pointer.
	2.	Establishes a new frame base.
	3.	Allocates local stack space.

Stack now looks like:

| local vars        |
| saved rbp         |
| return address    |

---

@[b005] pitfall
Common misconception:

“Function calls are just jumps.”

They are not.

A `jmp` does not push a return address.
A `call` modifies the stack and establishes structured return flow.


@[b006] checklist debugging
When debugging a crash inside a function:

- Is RSP aligned to 16 bytes?
- Is the return address valid?
- Was RBP preserved correctly?
- Did stack corruption overwrite the return address?
- Is the calling convention respected?

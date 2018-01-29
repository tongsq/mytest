	.file	"test.c"
	.text
	.globl	check_equal
	.type	check_equal, @function
check_equal:
.LFB11:
	.cfi_startproc
	cmpl	%esi, %edi
	sete	%al
	movzbl	%al, %eax
	ret
	.cfi_endproc
.LFE11:
	.size	check_equal, .-check_equal
	.section	.rodata.str1.1,"aMS",@progbits,1
.LC0:
	.string	"3 == 4 ? %d\n"
.LC1:
	.string	"3 == 3 ? %d\n"
	.text
	.globl	main
	.type	main, @function
main:
.LFB12:
	.cfi_startproc
	subq	$8, %rsp
	.cfi_def_cfa_offset 16
	movl	$0, %esi
	movl	$.LC0, %edi
	movl	$0, %eax
	call	printf
	movl	$1, %esi
	movl	$.LC1, %edi
	movl	$0, %eax
	call	printf
	addq	$8, %rsp
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE12:
	.size	main, .-main
	.ident	"GCC: (GNU) 4.8.5 20150623 (Red Hat 4.8.5-16)"
	.section	.note.GNU-stack,"",@progbits

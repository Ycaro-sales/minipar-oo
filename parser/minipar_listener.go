// Code generated from Minipar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Minipar
import "github.com/antlr4-go/antlr/v4"

// MiniparListener is a complete listener for a parse tree produced by MiniparParser.
type MiniparListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterClass_decl is called when entering the class_decl production.
	EnterClass_decl(c *Class_declContext)

	// EnterClass_member is called when entering the class_member production.
	EnterClass_member(c *Class_memberContext)

	// EnterMethod_decl is called when entering the method_decl production.
	EnterMethod_decl(c *Method_declContext)

	// EnterMethod_call is called when entering the method_call production.
	EnterMethod_call(c *Method_callContext)

	// EnterObj_creation is called when entering the obj_creation production.
	EnterObj_creation(c *Obj_creationContext)

	// EnterFunc_decl is called when entering the func_decl production.
	EnterFunc_decl(c *Func_declContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterVar_decl is called when entering the var_decl production.
	EnterVar_decl(c *Var_declContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterCompound_stmt is called when entering the compound_stmt production.
	EnterCompound_stmt(c *Compound_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterDo_stmt is called when entering the do_stmt production.
	EnterDo_stmt(c *Do_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterSeq_stmt is called when entering the seq_stmt production.
	EnterSeq_stmt(c *Seq_stmtContext)

	// EnterPar_stmt is called when entering the par_stmt production.
	EnterPar_stmt(c *Par_stmtContext)

	// EnterPrint_call is called when entering the print_call production.
	EnterPrint_call(c *Print_callContext)

	// EnterInput_call is called when entering the input_call production.
	EnterInput_call(c *Input_callContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterExpr_stmt is called when entering the expr_stmt production.
	EnterExpr_stmt(c *Expr_stmtContext)

	// EnterChannel_stmt is called when entering the channel_stmt production.
	EnterChannel_stmt(c *Channel_stmtContext)

	// EnterS_channel_stmt is called when entering the s_channel_stmt production.
	EnterS_channel_stmt(c *S_channel_stmtContext)

	// EnterC_channel_stmt is called when entering the c_channel_stmt production.
	EnterC_channel_stmt(c *C_channel_stmtContext)

	// EnterSend_stmt is called when entering the send_stmt production.
	EnterSend_stmt(c *Send_stmtContext)

	// EnterReceive_stmt is called when entering the receive_stmt production.
	EnterReceive_stmt(c *Receive_stmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOr_expr is called when entering the or_expr production.
	EnterOr_expr(c *Or_exprContext)

	// EnterAnd_expr is called when entering the and_expr production.
	EnterAnd_expr(c *And_exprContext)

	// EnterComparison_expr is called when entering the comparison_expr production.
	EnterComparison_expr(c *Comparison_exprContext)

	// EnterAdditive_expr is called when entering the additive_expr production.
	EnterAdditive_expr(c *Additive_exprContext)

	// EnterMultiplicative_expr is called when entering the multiplicative_expr production.
	EnterMultiplicative_expr(c *Multiplicative_exprContext)

	// EnterUnary_expr is called when entering the unary_expr production.
	EnterUnary_expr(c *Unary_exprContext)

	// EnterPostfix_expr is called when entering the postfix_expr production.
	EnterPostfix_expr(c *Postfix_exprContext)

	// EnterPrimary_expr is called when entering the primary_expr production.
	EnterPrimary_expr(c *Primary_exprContext)

	// EnterList_literal is called when entering the list_literal production.
	EnterList_literal(c *List_literalContext)

	// EnterDict_literal is called when entering the dict_literal production.
	EnterDict_literal(c *Dict_literalContext)

	// EnterKey_value_pair is called when entering the key_value_pair production.
	EnterKey_value_pair(c *Key_value_pairContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitClass_decl is called when exiting the class_decl production.
	ExitClass_decl(c *Class_declContext)

	// ExitClass_member is called when exiting the class_member production.
	ExitClass_member(c *Class_memberContext)

	// ExitMethod_decl is called when exiting the method_decl production.
	ExitMethod_decl(c *Method_declContext)

	// ExitMethod_call is called when exiting the method_call production.
	ExitMethod_call(c *Method_callContext)

	// ExitObj_creation is called when exiting the obj_creation production.
	ExitObj_creation(c *Obj_creationContext)

	// ExitFunc_decl is called when exiting the func_decl production.
	ExitFunc_decl(c *Func_declContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitVar_decl is called when exiting the var_decl production.
	ExitVar_decl(c *Var_declContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitCompound_stmt is called when exiting the compound_stmt production.
	ExitCompound_stmt(c *Compound_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitDo_stmt is called when exiting the do_stmt production.
	ExitDo_stmt(c *Do_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitSeq_stmt is called when exiting the seq_stmt production.
	ExitSeq_stmt(c *Seq_stmtContext)

	// ExitPar_stmt is called when exiting the par_stmt production.
	ExitPar_stmt(c *Par_stmtContext)

	// ExitPrint_call is called when exiting the print_call production.
	ExitPrint_call(c *Print_callContext)

	// ExitInput_call is called when exiting the input_call production.
	ExitInput_call(c *Input_callContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitExpr_stmt is called when exiting the expr_stmt production.
	ExitExpr_stmt(c *Expr_stmtContext)

	// ExitChannel_stmt is called when exiting the channel_stmt production.
	ExitChannel_stmt(c *Channel_stmtContext)

	// ExitS_channel_stmt is called when exiting the s_channel_stmt production.
	ExitS_channel_stmt(c *S_channel_stmtContext)

	// ExitC_channel_stmt is called when exiting the c_channel_stmt production.
	ExitC_channel_stmt(c *C_channel_stmtContext)

	// ExitSend_stmt is called when exiting the send_stmt production.
	ExitSend_stmt(c *Send_stmtContext)

	// ExitReceive_stmt is called when exiting the receive_stmt production.
	ExitReceive_stmt(c *Receive_stmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOr_expr is called when exiting the or_expr production.
	ExitOr_expr(c *Or_exprContext)

	// ExitAnd_expr is called when exiting the and_expr production.
	ExitAnd_expr(c *And_exprContext)

	// ExitComparison_expr is called when exiting the comparison_expr production.
	ExitComparison_expr(c *Comparison_exprContext)

	// ExitAdditive_expr is called when exiting the additive_expr production.
	ExitAdditive_expr(c *Additive_exprContext)

	// ExitMultiplicative_expr is called when exiting the multiplicative_expr production.
	ExitMultiplicative_expr(c *Multiplicative_exprContext)

	// ExitUnary_expr is called when exiting the unary_expr production.
	ExitUnary_expr(c *Unary_exprContext)

	// ExitPostfix_expr is called when exiting the postfix_expr production.
	ExitPostfix_expr(c *Postfix_exprContext)

	// ExitPrimary_expr is called when exiting the primary_expr production.
	ExitPrimary_expr(c *Primary_exprContext)

	// ExitList_literal is called when exiting the list_literal production.
	ExitList_literal(c *List_literalContext)

	// ExitDict_literal is called when exiting the dict_literal production.
	ExitDict_literal(c *Dict_literalContext)

	// ExitKey_value_pair is called when exiting the key_value_pair production.
	ExitKey_value_pair(c *Key_value_pairContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)
}

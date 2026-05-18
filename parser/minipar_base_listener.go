// Code generated from Minipar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Minipar
import "github.com/antlr4-go/antlr/v4"

// BaseMiniparListener is a complete listener for a parse tree produced by MiniparParser.
type BaseMiniparListener struct{}

var _ MiniparListener = &BaseMiniparListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMiniparListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMiniparListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMiniparListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMiniparListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMiniparListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMiniparListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseMiniparListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseMiniparListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterClass_decl is called when production class_decl is entered.
func (s *BaseMiniparListener) EnterClass_decl(ctx *Class_declContext) {}

// ExitClass_decl is called when production class_decl is exited.
func (s *BaseMiniparListener) ExitClass_decl(ctx *Class_declContext) {}

// EnterClass_member is called when production class_member is entered.
func (s *BaseMiniparListener) EnterClass_member(ctx *Class_memberContext) {}

// ExitClass_member is called when production class_member is exited.
func (s *BaseMiniparListener) ExitClass_member(ctx *Class_memberContext) {}

// EnterMethod_decl is called when production method_decl is entered.
func (s *BaseMiniparListener) EnterMethod_decl(ctx *Method_declContext) {}

// ExitMethod_decl is called when production method_decl is exited.
func (s *BaseMiniparListener) ExitMethod_decl(ctx *Method_declContext) {}

// EnterMethod_call is called when production method_call is entered.
func (s *BaseMiniparListener) EnterMethod_call(ctx *Method_callContext) {}

// ExitMethod_call is called when production method_call is exited.
func (s *BaseMiniparListener) ExitMethod_call(ctx *Method_callContext) {}

// EnterObj_creation is called when production obj_creation is entered.
func (s *BaseMiniparListener) EnterObj_creation(ctx *Obj_creationContext) {}

// ExitObj_creation is called when production obj_creation is exited.
func (s *BaseMiniparListener) ExitObj_creation(ctx *Obj_creationContext) {}

// EnterFunc_decl is called when production func_decl is entered.
func (s *BaseMiniparListener) EnterFunc_decl(ctx *Func_declContext) {}

// ExitFunc_decl is called when production func_decl is exited.
func (s *BaseMiniparListener) ExitFunc_decl(ctx *Func_declContext) {}

// EnterParams is called when production params is entered.
func (s *BaseMiniparListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseMiniparListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseMiniparListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseMiniparListener) ExitParam(ctx *ParamContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BaseMiniparListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BaseMiniparListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseMiniparListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseMiniparListener) ExitArgs(ctx *ArgsContext) {}

// EnterVar_decl is called when production var_decl is entered.
func (s *BaseMiniparListener) EnterVar_decl(ctx *Var_declContext) {}

// ExitVar_decl is called when production var_decl is exited.
func (s *BaseMiniparListener) ExitVar_decl(ctx *Var_declContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseMiniparListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseMiniparListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMiniparListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMiniparListener) ExitBlock(ctx *BlockContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseMiniparListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseMiniparListener) ExitStmt(ctx *StmtContext) {}

// EnterCompound_stmt is called when production compound_stmt is entered.
func (s *BaseMiniparListener) EnterCompound_stmt(ctx *Compound_stmtContext) {}

// ExitCompound_stmt is called when production compound_stmt is exited.
func (s *BaseMiniparListener) ExitCompound_stmt(ctx *Compound_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseMiniparListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseMiniparListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseMiniparListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseMiniparListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterDo_stmt is called when production do_stmt is entered.
func (s *BaseMiniparListener) EnterDo_stmt(ctx *Do_stmtContext) {}

// ExitDo_stmt is called when production do_stmt is exited.
func (s *BaseMiniparListener) ExitDo_stmt(ctx *Do_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseMiniparListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseMiniparListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterSeq_stmt is called when production seq_stmt is entered.
func (s *BaseMiniparListener) EnterSeq_stmt(ctx *Seq_stmtContext) {}

// ExitSeq_stmt is called when production seq_stmt is exited.
func (s *BaseMiniparListener) ExitSeq_stmt(ctx *Seq_stmtContext) {}

// EnterPar_stmt is called when production par_stmt is entered.
func (s *BaseMiniparListener) EnterPar_stmt(ctx *Par_stmtContext) {}

// ExitPar_stmt is called when production par_stmt is exited.
func (s *BaseMiniparListener) ExitPar_stmt(ctx *Par_stmtContext) {}

// EnterPrint_call is called when production print_call is entered.
func (s *BaseMiniparListener) EnterPrint_call(ctx *Print_callContext) {}

// ExitPrint_call is called when production print_call is exited.
func (s *BaseMiniparListener) ExitPrint_call(ctx *Print_callContext) {}

// EnterInput_call is called when production input_call is entered.
func (s *BaseMiniparListener) EnterInput_call(ctx *Input_callContext) {}

// ExitInput_call is called when production input_call is exited.
func (s *BaseMiniparListener) ExitInput_call(ctx *Input_callContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BaseMiniparListener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BaseMiniparListener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterExpr_stmt is called when production expr_stmt is entered.
func (s *BaseMiniparListener) EnterExpr_stmt(ctx *Expr_stmtContext) {}

// ExitExpr_stmt is called when production expr_stmt is exited.
func (s *BaseMiniparListener) ExitExpr_stmt(ctx *Expr_stmtContext) {}

// EnterChannel_stmt is called when production channel_stmt is entered.
func (s *BaseMiniparListener) EnterChannel_stmt(ctx *Channel_stmtContext) {}

// ExitChannel_stmt is called when production channel_stmt is exited.
func (s *BaseMiniparListener) ExitChannel_stmt(ctx *Channel_stmtContext) {}

// EnterS_channel_stmt is called when production s_channel_stmt is entered.
func (s *BaseMiniparListener) EnterS_channel_stmt(ctx *S_channel_stmtContext) {}

// ExitS_channel_stmt is called when production s_channel_stmt is exited.
func (s *BaseMiniparListener) ExitS_channel_stmt(ctx *S_channel_stmtContext) {}

// EnterC_channel_stmt is called when production c_channel_stmt is entered.
func (s *BaseMiniparListener) EnterC_channel_stmt(ctx *C_channel_stmtContext) {}

// ExitC_channel_stmt is called when production c_channel_stmt is exited.
func (s *BaseMiniparListener) ExitC_channel_stmt(ctx *C_channel_stmtContext) {}

// EnterSend_stmt is called when production send_stmt is entered.
func (s *BaseMiniparListener) EnterSend_stmt(ctx *Send_stmtContext) {}

// ExitSend_stmt is called when production send_stmt is exited.
func (s *BaseMiniparListener) ExitSend_stmt(ctx *Send_stmtContext) {}

// EnterReceive_stmt is called when production receive_stmt is entered.
func (s *BaseMiniparListener) EnterReceive_stmt(ctx *Receive_stmtContext) {}

// ExitReceive_stmt is called when production receive_stmt is exited.
func (s *BaseMiniparListener) ExitReceive_stmt(ctx *Receive_stmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMiniparListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMiniparListener) ExitExpr(ctx *ExprContext) {}

// EnterOr_expr is called when production or_expr is entered.
func (s *BaseMiniparListener) EnterOr_expr(ctx *Or_exprContext) {}

// ExitOr_expr is called when production or_expr is exited.
func (s *BaseMiniparListener) ExitOr_expr(ctx *Or_exprContext) {}

// EnterAnd_expr is called when production and_expr is entered.
func (s *BaseMiniparListener) EnterAnd_expr(ctx *And_exprContext) {}

// ExitAnd_expr is called when production and_expr is exited.
func (s *BaseMiniparListener) ExitAnd_expr(ctx *And_exprContext) {}

// EnterComparison_expr is called when production comparison_expr is entered.
func (s *BaseMiniparListener) EnterComparison_expr(ctx *Comparison_exprContext) {}

// ExitComparison_expr is called when production comparison_expr is exited.
func (s *BaseMiniparListener) ExitComparison_expr(ctx *Comparison_exprContext) {}

// EnterAdditive_expr is called when production additive_expr is entered.
func (s *BaseMiniparListener) EnterAdditive_expr(ctx *Additive_exprContext) {}

// ExitAdditive_expr is called when production additive_expr is exited.
func (s *BaseMiniparListener) ExitAdditive_expr(ctx *Additive_exprContext) {}

// EnterMultiplicative_expr is called when production multiplicative_expr is entered.
func (s *BaseMiniparListener) EnterMultiplicative_expr(ctx *Multiplicative_exprContext) {}

// ExitMultiplicative_expr is called when production multiplicative_expr is exited.
func (s *BaseMiniparListener) ExitMultiplicative_expr(ctx *Multiplicative_exprContext) {}

// EnterUnary_expr is called when production unary_expr is entered.
func (s *BaseMiniparListener) EnterUnary_expr(ctx *Unary_exprContext) {}

// ExitUnary_expr is called when production unary_expr is exited.
func (s *BaseMiniparListener) ExitUnary_expr(ctx *Unary_exprContext) {}

// EnterPostfix_expr is called when production postfix_expr is entered.
func (s *BaseMiniparListener) EnterPostfix_expr(ctx *Postfix_exprContext) {}

// ExitPostfix_expr is called when production postfix_expr is exited.
func (s *BaseMiniparListener) ExitPostfix_expr(ctx *Postfix_exprContext) {}

// EnterPrimary_expr is called when production primary_expr is entered.
func (s *BaseMiniparListener) EnterPrimary_expr(ctx *Primary_exprContext) {}

// ExitPrimary_expr is called when production primary_expr is exited.
func (s *BaseMiniparListener) ExitPrimary_expr(ctx *Primary_exprContext) {}

// EnterList_literal is called when production list_literal is entered.
func (s *BaseMiniparListener) EnterList_literal(ctx *List_literalContext) {}

// ExitList_literal is called when production list_literal is exited.
func (s *BaseMiniparListener) ExitList_literal(ctx *List_literalContext) {}

// EnterDict_literal is called when production dict_literal is entered.
func (s *BaseMiniparListener) EnterDict_literal(ctx *Dict_literalContext) {}

// ExitDict_literal is called when production dict_literal is exited.
func (s *BaseMiniparListener) ExitDict_literal(ctx *Dict_literalContext) {}

// EnterKey_value_pair is called when production key_value_pair is entered.
func (s *BaseMiniparListener) EnterKey_value_pair(ctx *Key_value_pairContext) {}

// ExitKey_value_pair is called when production key_value_pair is exited.
func (s *BaseMiniparListener) ExitKey_value_pair(ctx *Key_value_pairContext) {}

// EnterType is called when production type is entered.
func (s *BaseMiniparListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseMiniparListener) ExitType(ctx *TypeContext) {}

// EnterId is called when production id is entered.
func (s *BaseMiniparListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseMiniparListener) ExitId(ctx *IdContext) {}

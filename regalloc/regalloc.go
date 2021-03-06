package regalloc

import (
    . "g9cc/common"
    . "g9cc/util"
    . "g9cc/regs"
    . "g9cc/irdump"
)

// Register allocator
var used []bool = make([]bool, Len_Regs)

// IRの命令数分の要素をもつ配列(Alloc_regs()で初期化)
// 各中間表現にどのレジスタを割り当てるかを記録する配列
// 添字: 中間表現の順番, 要素: （レジスタの配列Regsの）何番目のレジスタを割り当てるか
var reg_map [8192]int

func alloc(ir_reg int) int {

    if reg_map[ir_reg] != -1 {
        var r int = reg_map[ir_reg]
        Assert(used[r], "allocation error")
        return r
    }

    // i はレジスタの配列regsのindex
    for i := 0; i < Len_Regs; i++ {
        // index i のレジスタが使用済みの場合
        if used[i] {
            continue
        }
        // index i のレジスタが未使用の場合
        reg_map[ir_reg] = i // registerへのmapping
        used[i] = true
        return i
    }

    Error("register exhausted")
    return 0 // ここには到達しないため(intを返さないと怒るコンパイラを鎮める他に)イミなし
}

func visit(irv *Vector) {
    // r0 is a reserved register that is always mapped to base pointer(rbp)
    reg_map[0] = 0
    used[0] = true

    for i := 0; i < irv.Len; i++ {
        ir := irv.Data[i].(*IR)

        switch Irinfo_arr[ir.Op].Ty {
        case IR_TY_BINARY:
            ir.Lhs = alloc(ir.Lhs)
            if !ir.Is_imm {
                ir.Rhs = alloc(ir.Rhs)
            }
        case IR_TY_REG, IR_TY_REG_IMM, IR_TY_REG_LABEL, IR_TY_LABEL_ADDR:
            ir.Lhs = alloc(ir.Lhs)
        case IR_TY_MEM, IR_TY_REG_REG:
            ir.Lhs = alloc(ir.Lhs)
            ir.Rhs = alloc(ir.Rhs)
        case IR_TY_CALL:
            ir.Lhs = alloc(ir.Lhs)
            for i := 0; i < ir.Nargs; i++ {
                ir.Args[i] = alloc(ir.Args[i])
            }
        }

        if ir.Op == IR_KILL {
            Assert(used[ir.Lhs], "kill error")
            used[ir.Lhs] = false
            ir.Op = IR_NOP
        }
    }
}

// 中間表現の命令配列fnsの各要素に対し、必要ならレジスタを割り当てていく
func Alloc_regs(fns *Vector) {
    for i := 0; i < len(reg_map); i++ {
        reg_map[i] = -1
    }
    //used = make([]bool, Len_Regs)

    for i := 0; i < fns.Len; i++ {
        fn := fns.Data[i].(*Function)

        // reg_map = make([]int, fn.Ir.Len)
        for j := 0; j < fn.Ir.Len; j++ {
            reg_map[j] = -1
        }
        //used = make([]bool, Len_Regs)

        visit(fn.Ir)
    }
}

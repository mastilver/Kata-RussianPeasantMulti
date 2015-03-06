package peasant_multiplication

import(
    "fmt"
)


func Multiply(left int, right int) int {
    var total int = 0;


    // return zero if one of the number is zero
    if left == 0 || right == 0 {
        return 0;
    }

    // protect against negative number
    var sign = 1;

    if left < 0 {
        left = -1 * left;
        sign = -1 * sign;
    }

    if right < 0 {
        right = -1 * right;
        sign = -1 * sign;
    }


    for {

        fmt.Printf("%d * %d\n", left, right);

        if left % 2 != 0 {
            total = total + right;
        }

        if(left == 1){
            break;
        }

        left = left / 2;
        right = right * 2;
    }

    return sign * total;
}

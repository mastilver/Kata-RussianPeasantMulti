package russian_peasant_multiplication

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
        left *= -1;
        sign = -1;
    }

    if right < 0 {
        right *= -1;
        sign *= -1;
    }


    for {

        fmt.Printf("%d * %d\n", left, right);

        if left % 2 != 0 {
            total += right;
        }

        if(left == 1){
            break;
        }

        left /= 2;
        right *= 2;
    }

    return sign * total;
}

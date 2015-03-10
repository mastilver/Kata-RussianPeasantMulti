package russian_peasant_multiplication

import(
    "fmt"
)


func Multiply(left int, right int) int {
    var total int = 0;
    var sign int;

    left, right, sign = CleanSign(left, right);

    for left != 0 {

        fmt.Printf("%d * %d\n", left, right);

        if left % 2 != 0 {
            total += right;
        }

        left /= 2;
        right *= 2;
    }

    return sign * total;
}

func RecurciveMultiply(left int, right int) int {

    var sign int;

    left, right, sign = CleanSign(left, right);

    if sign == 0 {
        return 0;
    } else if left % 2 == 0 {
        return sign * RecurciveMultiply(left / 2, right * 2);
    } else {
        return sign * (RecurciveMultiply(left / 2, right * 2) + right);
    }
}

func CleanSign(left int, right int) (int, int, int){

    if left == 0 || right == 0 {
        return left, right, 0;
    }

    var sign = 1;

    if left < 0 {
        left *= -1;
        sign = -1;
    }

    if right < 0 {
        right *= -1;
        sign *= -1;
    }

    return left, right, sign;
}

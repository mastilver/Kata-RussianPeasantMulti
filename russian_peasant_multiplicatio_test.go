package russian_peasant_multiplication

import (
    "math/rand"
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_Multiply_1(t *testing.T){
    assert := assert.New(t)

    assert.Equal(0, Multiply(0, 0), "0 * 0 should equal 0")
    assert.Equal(-100, Multiply(-5, 20), "-5 * 20 should equal -100")
    assert.Equal(-1337, Multiply(191, -7), "191 * -7 should equal -1337")
    assert.Equal(400, Multiply(20, 20), "-20 * -20 should equal 400")

    for i := 0; i < 10; i++ {

        var left = rand.Int();
        var right = rand.Int();
        var result = left * right;

        assert.Equal(result, Multiply(left, right), fmt.Sprintf("%d * %d should equal %d", left, right, result))
    }
}

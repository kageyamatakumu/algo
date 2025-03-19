// 問題文
// 次の 10 個の整数のうち、左から K 番目の整数を求めてください。
// 3,1,4,1,5,9,2,6,5,3
// ただし、最も左の整数を左から 0 番目と数えます
// K は 0 以上 9 以下の整数
var CHECK_ARRAY = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3];
var getKthElement = function (arr, index) {
    return arr[index];
};
console.log(getKthElement(CHECK_ARRAY, 3));
console.log(getKthElement(CHECK_ARRAY, 9));

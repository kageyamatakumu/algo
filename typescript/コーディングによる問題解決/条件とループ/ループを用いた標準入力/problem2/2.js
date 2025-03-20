// 問題文
// N 個の正の整数 A0,A1,…,AN-1が与えられます。N 個の整数を全て掛け合わせた値を求めてください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// Ai​は 1 以上 1000 以下の整数 (0≤i≤N−1)
var MIN_VALUE = 1;
var MAX_VALUE = 1000;
/** Nの数とAの要素数が一致しているか確認 */
var checkLength = function (N, A) {
    return N === A.length;
};
/** 1 以上 1000 以下の整数であるか確認 */
var checkNumberInRange = function (num) {
    return num >= MIN_VALUE && num <= MAX_VALUE;
};
/** 配列の要素が全て 1 以上 1000 以下の整数であるか確認 */
var checkArrayNumbersInRange = function (arr) {
    return arr.every(function (num) { return checkNumberInRange(num); });
};
/** 配列の合計を求める */
var multiplicationOfArray = function (arr) {
    return arr.reduce(function (num, element) { return num *= element; }, 1);
};
var calculateAndPrintResult = function (N, A) {
    if (!checkNumberInRange(N)) {
        return "N\u306F".concat(MIN_VALUE, "\u4EE5\u4E0A").concat(MAX_VALUE, "\u4EE5\u4E0B\u306E\u6574\u6570\u3092\u5165\u529B\u3057\u3066\u304F\u3060\u3055\u3044");
    }
    if (!checkLength(N, A)) {
        return "Nの値と配列の要素数が一致しません";
    }
    if (!checkArrayNumbersInRange(A)) {
        return "\u914D\u5217\u306E\u8981\u7D20\u306F".concat(MIN_VALUE, "\u4EE5\u4E0A").concat(MAX_VALUE, "\u4EE5\u4E0B\u306E\u6574\u6570\u3092\u5165\u529B\u3057\u3066\u304F\u3060\u3055\u3044");
    }
    return multiplicationOfArray(A);
};
console.log(calculateAndPrintResult(3, [10, 20, 30]));
console.log(calculateAndPrintResult(2, [10, 20, 30])); // エラーケース
console.log(calculateAndPrintResult(1001, [10, 20])); // エラーケース
console.log(calculateAndPrintResult(3, [0, 20, 30])); // エラーケース

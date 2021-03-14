package TypeSystem;

// 积类型 product type
// 取值为两字段的笛卡尔积
// final class Student {
//     String name;
//     int id;
// }

// 和类型 sum type
// 类型SchoolPerson可能是Student或者Teacher
// 类似取值范围，枚举
interface SchoolPerson {}
final class Student implements SchoolPerson {
    String name;
    int id;
}
final class Teacher implements SchoolPerson {
    String name;
    String office;
}

// 代数数据类型 ADT
// 由积类型 和类型组合构造出的类型
// 和类型具有枚举特性，积类型具有组合特性
interface Bool{}
final class True implements Bool {}// 可使用 instanceof 做判等谓词
final class False implements Bool {}
interface Nat {}// 自然数类型
final class Z implements Nat {}// 0类型
final class S implements Nat {// 由0构造其他自然数
    Nat value;
    S(Nat v) {value = v;}
}
// 3可以由 new S(new S(new S(new Z)))构造，意义同构
interface List<T> {}// 链表，广义表类型
final class Nil<T> implements List<T> {}// 递归基，平凡构造
final class Cons<T> implements List<T> {
    T value;
    List<T> next;

    Cons(T v, List<T> n) {
        value = v;
        next = n;
    }
}

// 适合构造树状结构，聚合数据结构
// 类似union，上述Bool类型，不同具体类型的树节点类型？

public class ADTPractice {
    public static void main(String[] args) {
        
    }
}
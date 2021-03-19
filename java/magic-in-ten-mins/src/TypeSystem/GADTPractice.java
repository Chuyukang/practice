package TypeSystem;

// GADT 广义代数数据类型
// 以类型作为参数的类型？

// Expr<Integer> = IVal + Add
// Expr<Boolean> = BVal + Eq<T>
// Add = Expr<Integer> * Expr<Integer>
// Eq<T> = Expr<T> * Expr<T>
interface Expr<T> {
    T getValue();
}
class IVal implements Expr<Integer> {
    Integer value;

    public IVal(int value) {
        this.value = value;
    }

    @Override
    public Integer getValue() {
        return value;
    }
}
class BVal implements Expr<Boolean> {
    Boolean value;

    public BVal(boolean value) {
        this.value = value;
    }

    @Override
    public Boolean getValue() {
        return value;
    }
}
class Add implements Expr<Integer> {
    Expr<Integer> e1,e2;
    public Add(Expr<Integer> a, Expr<Integer> b) {
        this.e1 = a;
        this.e2 = b;
    }

    @Override
    public Integer getValue() {
        return e1.getValue() + e2.getValue();
    }
}
class Eq<T> implements Expr<Boolean> {
    Expr<T> e1, e2;
    public Eq(Expr<T> e1, Expr<T> e2) {
        this.e1 = e1;
        this.e2 = e2;
    }
    @Override
    public Boolean getValue() {
        return e1.getValue() == e2.getValue();
    }
}


public class GADTPractice {
    public static void main(String[] args) {
        Expr<Integer> e1=new IVal(1), e2=new IVal(2);
        Expr<Integer> e3 = new Add(e1, e2);
        System.out.println(e3.getValue());

        Expr<Boolean> b1=new BVal(true), b2=new BVal(false);
        Expr<Boolean> b3=new Eq<>(b1, b2);
        System.out.println(b3.getValue());
    }

}

package src;

import frontend.CodeRunner;

public class run {
    public static void main(String[] args) {
        CodeRunner codeRunner = new CodeRunner();
        codeRunner.main("CodeRunner", 500, 400);
        codeRunner.frame();
    }
}

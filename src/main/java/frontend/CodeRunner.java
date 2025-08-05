package frontend;

import javax.swing.*;
import java.awt.*;

public class CodeRunner extends JFrame{

    public String title;
    public int width;
    public int height;

    public void main(String title, int width, int height) {
        this.title = title;
        this.width = width;
        this.height = height;
    }

    public void frame(){
        CodeRunner codeRunner = new CodeRunner();
        codeRunner.setTitle(this.title);
        codeRunner.setSize(this.width, this.height);
        codeRunner.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        codeRunner.setAlwaysOnTop(false);
        codeRunner.setLayout(null);
        codeRunner.setResizable(false);
        _items(codeRunner);
        codeRunner.setVisible(true);
    }

    private void _items(CodeRunner codeRunner){

        Font titleFont = new Font("CodeRunner", Font.BOLD, 18);

        JLabel title = new JLabel();
        title.setText("CodeRunner");
        title.setBounds(191, 1, 120, 80);
        title.setFont(titleFont);
        codeRunner.add(title);


    }

}
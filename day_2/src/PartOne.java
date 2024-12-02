import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class PartOne {

    static ArrayList<Integer> parseLevels(String line){
        String[] levelsStr = line.split(" ");
        ArrayList<Integer> out = new ArrayList<>();
        for (String level : levelsStr){
            out.add(Integer.parseInt(level));
        }
        return out;
    }

    static boolean isSafe(ArrayList<Integer> levels){
        int diff;
        int monotonicity = 0;
        int thisMonotonicity = 0;
        for (int i = 0; i < levels.size() - 1; i++){
            diff = levels.get(i) - levels.get(i+1);
            if (monotonicity == 0){
                monotonicity = diff > 0 ? 1 : -1;
            }
            thisMonotonicity = diff > 0 ? 1 : -1;
            if (thisMonotonicity != monotonicity){
                return false;
            }
            if (Math.abs(diff) < 1 || Math.abs(diff) > 3){
                return false;
            }
        }
        return true;
    }

    public static void main(){
        File inputFile = new File("input.txt");
        Scanner fileReader;
        int totalSafe = 0;
        try {
            fileReader = new Scanner(inputFile);
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        }
        while (fileReader.hasNextLine()){
            String line = fileReader.nextLine();
            ArrayList<Integer> levels = parseLevels(line);
            if (isSafe(levels)){
                totalSafe ++;
            }
        }
        System.out.println("Total safe: " + totalSafe);
    }
}

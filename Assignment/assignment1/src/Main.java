import java.util.Scanner; 
import java.util.ArrayList;
import java.io.BufferedReader;
import java.io.FileReader;
public class Main {
    
    //read the csv file 
    public static void main(String[] argus){
        String fileName = "yellow_tripdata_2009-01-15_1hour_clean.csv" ;
        ArrayList<TripRecord> records = readFile(fileName); 
        
        //ask user input the EPS and MinPts for DBscan 
        Scanner input = new Scanner(System.in); 
        System.out.println("Please specify cluster epsilon and minPts:"); 
        double eps = input.nextDouble(); 
        int minpts = input.nextInt(); 

        DBScan db = new DBScan(eps, minpts,records); //build cluster 

        for(Cluster c: db.getCluster()){
            c.printCluster(); 
        }

    }

    public static ArrayList<TripRecord> readFile(String nameOfFile){
        ArrayList<TripRecord> data = new ArrayList<>(); 
        try{//read file
            BufferedReader br = new BufferedReader(new FileReader(nameOfFile)); 
            //GET four element in order to create the tripRecord 
            String line; 
            while((line = br.readLine())!= null){
                String [] dataArray = line.split(","); 
                String pickupDateTime = dataArray[4]; 
                GPScoord pickupLocation = new GPScoord(Double.parseDouble(dataArray[8]),Double.parseDouble(dataArray[9])); 
                GPScoord dropoffLocation = new GPScoord(Double.parseDouble(dataArray[12]),Double.parseDouble(dataArray[13]));
                double tripDistance = Double.parseDouble(dataArray[7]);  
                data.add(new TripRecord(pickupDateTime, pickupLocation, dropoffLocation, tripDistance)); 

            }

        }catch(Exception e){
            e.printStackTrace();
        }
        return data;

    }


}

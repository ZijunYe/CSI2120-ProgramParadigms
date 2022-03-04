import java.util.Scanner; 
import java.util.ArrayList;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;

public class Main {
    //read the csv file 
    public static void main(String[] args){
        String fileName = "tripData.csv" ;
        ArrayList<TripRecord> records = readFile(fileName); 
        
        //ask user input the EPS and MinPts for DBscan 
           Scanner input = new Scanner(System.in);
            System.out.println("Please specify cluster epsilon and minPts:"); 
            double eps = input.nextDouble(); 
            int minPts = input.nextInt(); 
            
            //DBScan db = new DBScan(0.0001, 5,records);
            DBScan db = new DBScan(eps, minPts,records); //build cluster 
            for(Cluster c: db.getClusters()){
                c.printCluster(); 
            }
        

    }

    public static ArrayList<TripRecord> readFile(String nameOfFile){
        ArrayList<TripRecord> data = new ArrayList<>(); 
    
        try  {
            File file = new File(nameOfFile);
            FileReader f = new FileReader(file);
            BufferedReader br = new BufferedReader(f);
            String line = " ";
            String[] dataArray;
            line = br.readLine();
            while ((line = br.readLine()) != null) {
                dataArray = line.split(",");
                String pickup_DateTime = dataArray[4];
                GPScoord pickup_Location = new GPScoord(Double.parseDouble(dataArray[8]), Double.parseDouble(dataArray[9]));
                GPScoord dropoff_Location = new GPScoord(Double.parseDouble(dataArray[12]), Double.parseDouble(dataArray[13]));
                double trip_Distance = Double.parseDouble(dataArray[7]);
                data.add(new TripRecord(pickup_DateTime, pickup_Location, dropoff_Location, trip_Distance));
            }
            br.close();
        }catch(Exception e){
            e.printStackTrace();
        }
        return data;

    }


}

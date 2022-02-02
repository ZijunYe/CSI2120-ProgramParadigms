import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.util.ArrayList;

public class TaxiClusters {

    public static void main(String[] args) {
        String file = "DataSet.csv";
        double eps = 0.0001;
        int minPts = 5;
        ArrayList<TripRecord> tripData = read_file(file);

        DBScan dbScan = new DBScan(eps, minPts, tripData);
        for(Cluster c: dbScan.get_clusters()){
            c.print_cluster_data();
        }
    }

    private static ArrayList<TripRecord> read_file(String file_name) {
        ArrayList<TripRecord> tripRecord = new ArrayList<TripRecord>();
        try{
            File file = new File(file_name);
            FileReader fr = new FileReader(file);
            BufferedReader br = new BufferedReader(fr);
            String line = " ";
            String[] dataArr;
            line = br.readLine();
            while ((line = br.readLine()) != null) {
                dataArr = line.split(",");
                String pickup_DateTime = dataArr[4];
                GPScoord pickup_Location = new GPScoord(Double.parseDouble(dataArr[8]), Double.parseDouble(dataArr[9]));
                GPScoord dropoff_Location = new GPScoord(Double.parseDouble(dataArr[12]), Double.parseDouble(dataArr[13]));
                double trip_Distance = Double.parseDouble(dataArr[7]);
                tripRecord.add(new TripRecord(pickup_DateTime, pickup_Location, dropoff_Location, trip_Distance));
            }
            br.close();
        }catch(Exception e){
            e.printStackTrace();
        }
        return tripRecord;
    }
}

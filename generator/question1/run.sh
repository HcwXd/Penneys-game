seq_len=3
head_prob=0.1

for seq_len in {3..6};
do
    for head_prob in {5..9};
    do
        head_prob="0.$head_prob"
        file_name="len_"
        file_name+=$seq_len
        file_name+="_headProb_"
        file_name+=$head_prob
        file_name+="_result.txt"
        echo $file_name
        ./main $seq_len $head_prob > $file_name
    done
done
